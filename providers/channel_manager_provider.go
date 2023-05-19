package providers

import (
	"board/config"
	"board/services/pusher"
	"fmt"
)

func ChannelManager(c config.Config) *pusher.ChannelManager {
	channelManager := pusher.NewChannelManager(c)

	go channelManager.Reactor()

	fmt.Println("init ChannelManager")

	return channelManager
}
