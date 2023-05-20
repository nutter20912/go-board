package pusher

import (
	"board/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

type Channel struct {
	config  config.Config
	Name    string
	Clients map[*Client]bool
}

func (c *Channel) subscribe(message *Message) (err error) {
	if err := c.verifySignature(message); err != nil {
		return err
	}

	c.Clients[message.client] = true

	return err
}

func (c *Channel) verifySignature(message *Message) (err error) {
	signature := message.client.socketId + ":" + c.Name
	hashed := generateHMAC(signature, c.config.Pusher.Secret)

	if strings.Split(message.Data.Auth, ":")[1] != hashed {
		return errors.New("auth fail")
	}

	return err
}

func generateHMAC(signature, secret string) string {
	key := []byte(secret)
	message := []byte(signature)

	h := hmac.New(sha256.New, key)
	h.Write(message)
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
