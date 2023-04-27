package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
}

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
