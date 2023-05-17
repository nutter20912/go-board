package controllers

import (
	"board/libs"
	"board/services/pusher"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PusherAction struct {
	Controller
}

func (p PusherAction) Sub(ctx *gin.Context) {
	channelManager := ctx.MustGet("channelManager").(*pusher.ChannelManager)

	conn, err := libs.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := pusher.Client{
		ChannelManager: channelManager,
		Conn:           conn,
		Send:           make(chan []byte, 256),
	}

	go client.ReadPump()
	go client.WritePump()

	client.Open()
}
