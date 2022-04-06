package middleware_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"

	"go-framework/internal/middleware"
	"go-framework/pkg/test"
)

func TestAuth(t *testing.T) {
	token, err := booted.JWT.Create(1)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	t.Run("无 token", func(t *testing.T) {
		cc, rec := test.NewMiddlewareTest()
		cc.JWT = booted.JWT
		cc.Request.Header.Set("Authorization", "")
		middleware.Auth()(cc.Context)
		assert.True(t, cc.IsAborted())
		assertCodeAndMsg(t, rec, http.StatusUnauthorized, middleware.NoToken)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("无效的 token", func(t *testing.T) {
		cc, rec := test.NewMiddlewareTest()
		cc.JWT = booted.JWT
		cc.Request.Header.Set("Authorization", "123456")
		middleware.Auth()(cc.Context)
		assert.True(t, cc.IsAborted())
		assertCodeAndMsg(t, rec, http.StatusUnauthorized, middleware.InvalidToken)
	})

	t.Run("有效的 token", func(t *testing.T) {
		cc, rec := test.NewMiddlewareTest()
		cc.JWT = booted.JWT
		cc.Request.Header.Set("Authorization", token.Token)
		middleware.Auth()(cc.Context)
		assert.False(t, cc.IsAborted(), fmt.Sprintf("want no aborted, but got aborted: %d, %s", rec.Code, rec.Body.String()))
		assert.Equal(t, uint64(1), cc.UserID)
	})
}

// 断言响应 code 与 message
func assertCodeAndMsg(t *testing.T, rec *test.ResponseRecorder, wantCode int, wantMsg string) bool {
	t.Helper()
	if rec == nil {
		t.Fatal("rec is nil")
	}
	if wantCode != rec.Code {
		t.Errorf("want code %d, got %d", wantCode, rec.Code)
		return false
	}
	msg := struct {
		Message string `json:"message"`
	}{}
	if !assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &msg)) {
		t.Fatal("unmarshal error")
	}
	if !strings.Contains(msg.Message, wantMsg) {
		t.Errorf("want msg contains %s, got %s", wantMsg, msg.Message)
		return false
	}
	return true
}
