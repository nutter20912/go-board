package pusher

type Data struct {
	Auth     string `json:"auth,omitempty"`
	Channel  string `json:"channel,omitempty"`
	Message  string `json:"message,omitempty"`
	Code     string `json:"code,omitempty"`
	SocketId string `json:"socket_id,omitempty"`
}

type Message struct {
	client  *Client
	Event   string `json:"event"`
	Channel string `json:"channel,omitempty"`
	*Data   `json:"data,omitempty"`
}

type ErrorMessage struct {
	Event string `json:"event"`
	Data  `json:"data"`
}

const (
	EVENT_CONNECTION_ESTABLISHED string = "pusher:connection_established"
	EVENT_SIGNIN                 string = "pusher:signin"
	EVENT_SIGNIN_SUCCESS         string = "pusher:signin_success"
	EVENT_ERROR                  string = "pusher:error"
	EVENT_SUBSCRIBE              string = "pusher:subscribe"
	EVENT_UNSUBSCRIBE            string = "pusher:unsubscribe"
	EVENT_SUBSCRIBE_SUCCESS      string = "pusher_internal:subscription_succeeded"
)
