package test

import (
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-framework/pkg/lgo"
)

type ResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *ResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func (r *ResponseRecorder) closeClient() {
	r.closeChannel <- true
}

func CreateTestResponseRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

// NewUploaderContext 创建上下文
//
// 该 ctx 已经经过 middlewares.WithUploaderContext 中间件处理
// 可以通过 uc.Request = http.NewRequest() 来更改请求对象，所需依赖均已注入
// 可以通过 uc.Write 来获取响应对象
func NewUploaderContext(dep *lgo.Dependencies) *lgo.CustomContext {
	rec := CreateTestResponseRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	lgo.WithContext(dep)(c)
	return c.MustGet(lgo.CustomContextKey).(*lgo.CustomContext)
}

// NewMiddlewareTest 创建中间件测试
// rec 用于测试中间件 abort
// 未注入不需要的依赖，中间件应当在创建时手动传入依赖
func NewMiddlewareTest() (uc *lgo.CustomContext, rec *ResponseRecorder) {
	rec = CreateTestResponseRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	uc = &lgo.CustomContext{Context: c}
	c.Set(lgo.CustomContextKey, uc)
	return uc, rec
}

// NewHTTPTest 创建 http 测试需要的 req、resp
func NewHTTPTest(method, target string, body io.Reader) (*http.Request, *ResponseRecorder) {
	req := httptest.NewRequest(method, target, body)
	rec := CreateTestResponseRecorder()
	return req, rec
}

func RandUserID() int64 {
	return NewRand().Int63()
}

func RandIP() net.IP {
	newRand := NewRand()
	return net.IPv4(byte(newRand.Intn(255)), byte(newRand.Intn(255)), byte(newRand.Intn(255)), byte(newRand.Intn(255)))
}

func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PrintRec(t *testing.T, rec *ResponseRecorder) {
	t.Logf("rec: code=%d,body=%s", rec.Code, rec.Body.String())
}
