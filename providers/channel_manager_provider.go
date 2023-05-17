package providers

import (
	"board/services/pusher"
	"fmt"
)

func ChannelManager() *pusher.ChannelManager {
	channelManager := pusher.NewChannelManager()

	go func() {
		for {
			select {
			case channel := <-channelManager.Register:
				fmt.Printf("register: %v - %v\n", channel.Name, channelManager.Channels)
				channelManager.Channels[channel.Name] = channel
			}
		}
	}()

	fmt.Println("init ChannelManager")

	return channelManager
}
