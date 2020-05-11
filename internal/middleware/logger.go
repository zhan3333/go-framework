package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 记录请求与响应的数据
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		startTime := time.Now()

		var request map[string]interface{}

		b, _ := c.Copy().GetRawData()

		_ = json.Unmarshal(b, &request)

		s, _ := json.Marshal(request)

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

		//处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		//结束时间
		endTime := time.Now()

		log.WithFields(log.Fields{
			"request_uri":    c.Request.RequestURI,
			"request_method": c.Request.Method,
			"client_ip":      c.ClientIP(),
			"request_time":   startTime.String(),
			"response_time":  endTime.String(),
			"request":        string(s),
			"response":       responseBody,
			"use_time":       endTime.Sub(startTime).String(),
		}).Info("记录请求")
	}
}
