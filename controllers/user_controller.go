package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserAction struct {
	Controller
}

type UserInput struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=3,max=20"`
}

// 新增 user
func (u UserAction) Store(ctx *gin.Context) {
	var input UserInput

	if err := u.validate(ctx, &input); err != nil {
		return
	}

	u.success(ctx, 201, gin.H{
		"user": input,
	})
}
