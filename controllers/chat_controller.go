package controllers

import (
	"board/libs"
	"board/services/chat"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatAction struct {
	Controller
}

// 聊天室
func (c ChatAction) Room(ctx *gin.Context) {
	conn, err := libs.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		return nil
	})

	hub := ctx.MustGet("hub").(*chat.Hub)

	client := &chat.Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
