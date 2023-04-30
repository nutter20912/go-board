package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 請求日誌
func RequestLogger(ctx *gin.Context) {
	log := ctx.MustGet("log").(*logrus.Logger)
	var requestBody map[string]interface{}
	var responseBody map[string]interface{}

	start := time.Now()
	bodyLogWriter := &bodyLogWriter{
		ResponseWriter: ctx.Writer,
		body:           bytes.NewBufferString(""),
	}
	ctx.Writer = bodyLogWriter

	ctx.Next()

	end := time.Now()

	body := ctx.Request.Body
	defer body.Close()

	bodyData, _ := ioutil.ReadAll(body)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
	json.Unmarshal(bodyData, &requestBody)
	json.Unmarshal(bodyLogWriter.body.Bytes(), &responseBody)

	log.
		WithFields(logrus.Fields{
			"latency": fmt.Sprint(end.Sub(start)),
			"ip":      ctx.ClientIP(),
			"method":  ctx.Request.Method,
			"request": logrus.Fields{
				"query": ctx.Request.URL.Query(),
				"body":  requestBody,
			},
			"response": logrus.Fields{
				"code": bodyLogWriter.ResponseWriter.Status(),
				"body": responseBody,
			},
		}).
		Info("requestLog")
}
