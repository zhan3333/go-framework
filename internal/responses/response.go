package responses

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson/buffer"
	validator2 "go-framework/internal/validator"
	"gopkg.in/go-playground/validator.v9"
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
	JSON(c, CodeFailed, err.Error(), nil)
}

// 请求参数不合法
func BadReq(c *gin.Context, err error) {
	var errMsg buffer.Buffer
	errData := err.(validator.ValidationErrors).Translate(validator2.Trans)
	// 取第一个错误信息
	for _, value := range errData {
		errMsg.AppendString(value)
		break
	}
	JSON(c, CodeFailed, string(errMsg.Buf), nil)
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
