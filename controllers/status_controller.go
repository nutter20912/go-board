package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusAction struct {
	Controller
}

func (s StatusAction) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
