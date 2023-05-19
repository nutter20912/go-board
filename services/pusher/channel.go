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

func (c *Channel) subscribe(protocolMessage *ProtocolMessage) (err error) {
	if err := c.verifySignature(protocolMessage); err != nil {
		return err
	}

	c.Clients[protocolMessage.client] = true

	return err
}

func (c *Channel) verifySignature(protocolMessage *ProtocolMessage) (err error) {
	signature := protocolMessage.client.socketId + ":" + c.Name
	hashed := generateHMAC(signature, c.config.Pusher.Secret)

	if strings.Split(protocolMessage.Data.Auth, ":")[1] != hashed {
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
