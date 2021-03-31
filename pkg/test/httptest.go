package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
)

type Http struct {
	Router *gin.Engine
}

func New(r *gin.Engine) *Http {
	return &Http{Router: r}
}

type Header struct {
	Key   string
	Value string
}

func (h *Http) Call(method, path string, body interface{}, headers ...Header) *httptest.ResponseRecorder {
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(method, path, strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	h.Router.ServeHTTP(w, req)
	return w
}

// Get 方法在绑定请求体的时候, 需要使用 ShouldBindJSON 方法 (gin 写法上限制的)
func (h *Http) Get(path string, body interface{}, headers ...Header) *httptest.ResponseRecorder {
	u, _ := url.Parse(path)
	values := u.Query()
	for k, v := range h.toMap(body) {
		values.Add(k, h.interface2String(v))
	}
	u.RawQuery = values.Encode()
	return h.Call(http.MethodGet, u.String(), nil, headers...)
}

func (h *Http) Post(path string, body interface{}, headers ...Header) *httptest.ResponseRecorder {
	return h.Call(http.MethodPost, path, body, headers...)
}

func (h *Http) toMap(body interface{}) map[string]interface{} {
	b, _ := json.Marshal(body)
	m := map[string]interface{}{}
	_ = json.Unmarshal(b, &m)
	return m
}

func (h *Http) interface2String(inter interface{}) string {

	switch inter.(type) {

	case string:
		return inter.(string)
	case int:
		return strconv.Itoa(inter.(int))
	case float64:
		return strconv.Itoa(int(inter.(float64)))
	}
	return ""
}
