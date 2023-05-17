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
	conn, err := libs.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := pusher.Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	go client.WritePump()

	client.Open()
}
