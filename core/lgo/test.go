package lgo

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Test struct {
	Router *gin.Engine
}

func New(r *gin.Engine) *Test {
	return &Test{Router: r}
}

func (h *Test) Send(req *http.Request) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	h.Router.ServeHTTP(resp, req)
	return resp
}

func ToReader(body interface{}) io.Reader {
	b, _ := json.Marshal(body)
	return strings.NewReader(string(b))
}
