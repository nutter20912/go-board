package pusher

import (
	"board/config"
)

type ChannelManager struct {
	config   config.Config
	Channels map[string]*Channel
	Register chan *ProtocolMessage
}

func NewChannelManager(c config.Config) *ChannelManager {
	return &ChannelManager{
		config:   c,
		Channels: make(map[string]*Channel),
		Register: make(chan *ProtocolMessage),
	}
}

// 事件處理
func (c *ChannelManager) Reactor() {
	for {
		select {
		case protocolMessage := <-c.Register:
			channel := c.findOrCreate(protocolMessage.Data.Channel)

			if err := channel.subscribe(protocolMessage); err != nil {
				protocolMessage.client.Send <- protocolMessage.client.jsonEncode(
					ErrorMessage{
						Event: EVENT_ERROR,
						Data:  Data{Message: err.Error()},
					})
			}
		}
	}
}

func (c *ChannelManager) findOrCreate(channelName string) *Channel {
	if channel, ok := c.Channels[channelName]; ok {
		return channel
	}

	channel := &Channel{
		config:  c.config,
		Name:    channelName,
		Clients: make(map[*Client]bool),
	}

	c.Channels[channelName] = channel

	return channel
}
