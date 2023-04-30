package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Method string
	Path   string
	Action gin.HandlerFunc
}

// 驗證器
func (c *Controller) validate(ctx *gin.Context, input interface{}) error {
	if err := ctx.ShouldBind(input); err != nil {
		c.error(ctx, http.StatusBadRequest, err.Error())

		return err
	}

	return nil
}

func (c *Controller) error(ctx *gin.Context, errorCode int, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    errorCode,
		"message": message,
		"result":  nil,
	})
}

func (c *Controller) success(ctx *gin.Context, code int, result gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"result":  result,
	})
}
