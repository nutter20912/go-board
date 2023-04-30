package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 全局錯誤處理
func ErrorHandler(ctx *gin.Context) {
	log := ctx.MustGet("log").(*logrus.Logger)

	defer func() {
		if err := recover(); err != nil {
			log.WithFields(logrus.Fields{
				"type":   "error",
				"method": ctx.Request.Method,
			}).Error(err)

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Internal Server Error",
			})
		}
	}()

	ctx.Next()
}
