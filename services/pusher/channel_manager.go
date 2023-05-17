package pusher

type Channel struct {
	Clients map[*Client]bool
}

type ChannleManager struct {
	Channels map[string]*Channel
}
