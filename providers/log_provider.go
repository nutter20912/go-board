package providers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}

	fmt.Println("init logger")

	return log
}
