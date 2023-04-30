package providers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(c *gin.Context) {
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}
	c.Set("log", log)

	c.Next()
}
