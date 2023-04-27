package bootstrap

import (
	"board/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	router := gin.Default()
	router.GET("/", controllers.Status)

	return &http.Server{
		Addr:    os.Getenv("APP_HOST") + os.Getenv("APP_PORT"),
		Handler: router,
	}
}
