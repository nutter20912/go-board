package pusher

import (
	"board/helper"
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

		var m = Message{}
		json.Unmarshal(message, &m)

		switch m.Event {
		case EVENT_SUBSCRIBE:
			c.subcribe(m)
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

func (c *Client) Open() {
	c.generateSocketId()
	c.onConnectioned()
}

// 生成 socket 編號
func (c *Client) generateSocketId() *Client {
	c.socketId = fmt.Sprintf("%d.%d", rand.Uint64(), rand.Uint64())

	return c
}

// 訂閱頻道
func (c *Client) subcribe(m Message) {
	m.client = c

	c.ChannelManager.Register <- &m
}

// 已建立連接
func (c *Client) onConnectioned() {
	message := Message{
		Event: EVENT_CONNECTION_ESTABLISHED,
		Data:  &Data{SocketId: c.socketId},
	}

	c.Send <- helper.JsonEncode(message)
}

// 訂閱成功
func (c *Client) onSubscribeSucceeded(channel string) {
	message := Message{
		Event:   EVENT_SUBSCRIBE_SUCCESS,
		Channel: channel,
	}

	c.Send <- helper.JsonEncode(message)
}

// 錯誤訊息
func (c *Client) onError(errMsg string) {
	message := ErrorMessage{
		Event: EVENT_ERROR,
		Data:  Data{Message: errMsg},
	}

	c.Send <- helper.JsonEncode(message)
}
