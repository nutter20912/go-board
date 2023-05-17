package chat

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

var (
	newline      = []byte{'\n'}
	space        = []byte{' '}
	PingPeriod   = time.Minute
	ReadTimeout  = PingPeriod + time.Second*5
	WriteTimeout = time.Second * 10
)

// 讀取訊息
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}

// 寫入流處理
func (c *Client) WritePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.writeHandler(websocket.CloseMessage, []byte{})
				return
			}
			c.writeHandler(websocket.TextMessage, []byte(message))
		case <-ticker.C:
			if err := c.writeHandler(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 處理寫入訊息
func (c *Client) writeHandler(messageType int, data []byte) error {
	c.Conn.SetWriteDeadline(time.Now().Add(WriteTimeout))

	return c.Conn.WriteMessage(messageType, data)
}
