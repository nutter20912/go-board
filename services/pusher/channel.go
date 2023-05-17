package pusher

type Channel struct {
	Name    string
	Clients map[*Client]bool
}

func (channel *Channel) subscribe(client *Client) {
	channel.Clients[client] = true
}

func (c *Channel) verifySignature() {

}
