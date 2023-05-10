package providers

import (
	"board/services/chat"
	"fmt"
)

func Hub() *chat.Hub {
	hub := chat.NewHub()

	go func() {
		for {
			select {
			case client := <-hub.Register:
				fmt.Printf("register: %v\n", hub.Clients)
				hub.Clients[client] = true
			case client := <-hub.Unregister:
				fmt.Printf("unregister: %v\n", hub.Clients)
				if _, ok := hub.Clients[client]; ok {
					delete(hub.Clients, client)
					close(client.Send)
				}
			case message := <-hub.Broadcast:
				fmt.Printf("message: %v\n", string(message))
				for client := range hub.Clients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(hub.Clients, client)
					}
				}
			}
		}
	}()

	fmt.Println("init chat hub")

	return hub
}
