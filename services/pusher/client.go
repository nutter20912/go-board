package pusher

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	socketId string
	Conn     *websocket.Conn
	Send     chan []byte
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
