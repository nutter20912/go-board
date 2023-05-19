package pusher

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ChannelManager *ChannelManager
	socketId       string
	Conn           *websocket.Conn
	Send           chan []byte
}

func (c *Client) jsonEncode(value interface{}) []byte {
	m, _ := json.Marshal(value)

	return m
}

func (c *Client) Open() {
	c.generateSocketId()

	c.Send <- c.jsonEncode(ProtocolMessage{
		Event: EVENT_CONNECTION_ESTABLISHED,
		Data:  Data{SocketId: c.socketId},
	})
}

func (c *Client) generateSocketId() *Client {
	c.socketId = fmt.Sprintf("%d.%d", rand.Uint64(), rand.Uint64())

	return c
}

// 讀取客戶端消息
func (c *Client) ReadPump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			break
		}

		var mes = ProtocolMessage{}
		json.Unmarshal(message, &mes)

		switch mes.Event {
		case "pusher:subscribe":
			c.ChannelManager.Register <- &ProtocolMessage{
				client: c,
				Event:  mes.Event,
				Data:   mes.Data,
			}
		}
	}
}

// 寫入客戶端消息
func (c *Client) WritePump() {
	ticker := time.NewTicker(time.Second * 10)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, []byte(message))

		case <-ticker.C:
			fmt.Println("ping")
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
