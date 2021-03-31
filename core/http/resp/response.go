package resp

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const CodeSuccess = 0
const CodeFailed = 1
const SuccessMsg = "success"
const FailMsg = "failed"

type ResponseStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Body    interface{} `json:"body"`
}

func Parse(bytes []byte) ResponseStruct {
	var response ResponseStruct
	_ = json.Unmarshal(bytes, &response)
	return response
}

type Responser interface {
	JSON(code int, msg string, body interface{})
	Success(msg string, body interface{})
	SuccessMsg(msg string)
	SuccessBody(body interface{})
	SuccessEmpty()
	Failed(msg string, body interface{})
	FailedMsg(msg string)
	FailedBody(body interface{})
	FailedEmpty()
	Error(err error, msg string, body interface{})
	ErrorWithBody(err error, body interface{})
	ErrorWithMsg(err error, msg string)
	ErrorEmpty(err error)
	// 绑定请求参数
	MustBind(obj interface{})
}

type Resp struct {
	c *gin.Context
}

func (r Resp) Success(msg string, body interface{}) {
	r.JSON(CodeSuccess, msg, body)
}

func (r Resp) SuccessMsg(msg string) {
	r.Success(msg, nil)
}

func (r Resp) SuccessBody(body interface{}) {
	r.Success(SuccessMsg, body)
}

func (r Resp) SuccessEmpty() {
	r.Success(SuccessMsg, nil)
}

func (r Resp) Failed(msg string, body interface{}) {
	r.JSON(CodeFailed, msg, body)
}

func (r Resp) FailedMsg(msg string) {
	r.Failed(msg, nil)
}

func (r Resp) FailedBody(body interface{}) {
	r.Failed(FailMsg, body)
}

func (r Resp) FailedEmpty() {
	r.Failed(FailMsg, nil)
}

func (r Resp) Error(err error, msg string, body interface{}) {
	r.JSON(CodeFailed, fmt.Sprintf("%s: %s", msg, err.Error()), body)
}

func (r Resp) ErrorWithBody(err error, body interface{}) {
	r.Error(err, FailMsg, body)
}

func (r Resp) ErrorWithMsg(err error, msg string) {
	r.Error(err, msg, nil)
}

func (r Resp) ErrorEmpty(err error) {
	r.Error(err, FailMsg, nil)
}

func (r Resp) MustBind(obj interface{}) {
	if err := r.c.ShouldBind(obj); err != nil {
		panic(err)
	}
}

func NewResp(c *gin.Context) Responser {
	return Resp{c: c}
}

func (r Resp) JSON(code int, msg string, body interface{}) {
	r.c.JSON(http.StatusOK, ResponseStruct{
		Code:    code,
		Message: msg,
		Body:    body,
	})
}
