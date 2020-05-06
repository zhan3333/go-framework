package Response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const CodeSuccess = 0
const CodeFailed = 1

func JSON(c *gin.Context, code int, message string, body interface{}) {
	c.JSON(http.StatusOK, struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}{
		Code:    code,
		Message: message,
		Body:    body,
	})
}

func Success(c *gin.Context, message string, body interface{}) {
	JSON(c, CodeSuccess, message, body)
}

func Failed(c *gin.Context, message string, body interface{}) {
	JSON(c, CodeFailed, message, body)
}

func Error(c *gin.Context, err error) {
	log.Errorf("Response error: %s", err.Error())
	JSON(c, CodeFailed, err.Error(), nil)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Parse(bytes []byte) Response {
	var response Response
	_ = json.Unmarshal(bytes, &response)
	return response
}
