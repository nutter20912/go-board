package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatAction struct {
	Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte("byebye"))
			}
		}
	}
}

func newHub() *Hub {
	hub := &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	go func() {
		for {
			select {
			case client := <-hub.register:
				hub.clients[client] = true
			case client := <-hub.unregister:
				if _, ok := hub.clients[client]; ok {
					delete(hub.clients, client)
					close(client.send)
				}
				//case message := <-hub.broadcast:
				//	for client := range hub.clients {
				//		select {
				//		case client.send <- message:
				//		default:
				//			close(client.send)
				//			delete(hub.clients, client)
				//		}
				//	}
			}
		}
	}()

	return hub
}
func (h *Hub) run() {

}

var hub = newHub()

// 聊天室
func (c ChatAction) Room(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		c.error(ctx, 400, "conn err")
		return
	}
	defer conn.Close()

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	go client.writePump()

	//for {
	//	conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format(time.RFC3339)))
	//	time.Sleep(time.Second)
	//}
}
