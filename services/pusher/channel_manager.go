package pusher

import (
	"board/config"
)

type ChannelManager struct {
	config   config.Config
	Channels map[string]*Channel
	Register chan *Message
}

func NewChannelManager(c config.Config) *ChannelManager {
	return &ChannelManager{
		config:   c,
		Channels: make(map[string]*Channel),
		Register: make(chan *Message),
	}
}

// 事件處理jsonEncode
func (cm *ChannelManager) Reactor() {
	for {
		select {
		case message := <-cm.Register:
			channel := cm.findOrCreate(message.Data.Channel)

			if err := channel.subscribe(message); err != nil {
				message.client.onError(err.Error())
				continue
			}

			message.client.onSubscribeSucceeded(message.Data.Channel)
		}
	}
}

func (cm *ChannelManager) findOrCreate(channelName string) *Channel {
	if channel, ok := cm.Channels[channelName]; ok {
		return channel
	}

	channel := &Channel{
		config:  cm.config,
		Name:    channelName,
		Clients: make(map[*Client]bool),
	}

	cm.Channels[channelName] = channel

	return channel
}
