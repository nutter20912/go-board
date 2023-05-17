package pusher

type ChannelManager struct {
	Channels map[string]*Channel
	Register chan *Channel
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		Channels: make(map[string]*Channel),
		Register: make(chan *Channel),
	}
}

func (c *ChannelManager) findOrCreate(channelName string) *Channel {
	if channel, ok := c.Channels[channelName]; ok {
		return channel
	}

	channel := &Channel{
		Name:    channelName,
		Clients: make(map[*Client]bool),
	}

	c.Channels[channelName] = channel

	return channel
}
