package controllers

import (
	"board/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserAction struct {
	Controller
}

type UserInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=3,max=20"`
}

// 查詢使用者
func (u UserAction) Show(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User

	id := ctx.Param("id")

	if err := db.First(&user, id).Error; err != nil {
		u.error(ctx, 400, fmt.Sprint(err))
		return
	}

	u.success(ctx, 200, gin.H{
		"user": user,
	})
}

// 新增使用者
func (u UserAction) Store(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var input UserInput
	var user models.User

	if err := u.validate(ctx, &input); err != nil {
		return
	}

	if rows := db.Where("email = ?", input.Email).First(&user, input.ID).RowsAffected; rows > 0 {
		u.error(ctx, 400, fmt.Sprint("重複email"))
		return
	}

	user = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		u.error(ctx, 401, fmt.Sprint(err))
		return
	}

	u.success(ctx, 201, gin.H{
		"user": user,
	})
}
