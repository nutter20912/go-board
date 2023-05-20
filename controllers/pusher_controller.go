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

type TriggerEventMessage struct {
	Name     string
	Data     interface{}
	Channel  string
	Channels []string
	SocketId string
}

// 觸發事件
func (p *PusherAction) TriggerEvent(ctx *gin.Context) {
	appId := ctx.Param("app_id")

	// ValidSignature
	// get channels
	// find channel
	// broadcast expect self

	p.success(ctx, 200, gin.H{
		"user": appId,
	})
}
