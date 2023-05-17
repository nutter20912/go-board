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

type Message struct {
	Event string            `json:"event"`
	Data  map[string]string `json:"data"`
}

func (c *Client) Open() {
	c.generateSocketId()

	m, _ := json.Marshal(Message{
		Event: "pusher:connection_established",
		Data: map[string]string{
			"socket_id": c.socketId,
		},
	})

	c.Send <- m
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

		var mes = Message{}
		json.Unmarshal(message, &mes)

		switch mes.Event {
		case "pusher:subscribe":
			//TODO 改成chan處理
			channel := c.ChannelManager.findOrCreate(mes.Data["channel"])
			channel.subscribe(c)
			fmt.Println(c.ChannelManager.Channels)
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
