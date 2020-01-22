package Controllers_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-framework/Tool"
	"go-framework/app/Http/Request"
	"go-framework/app/Http/Response"
	"go-framework/bootstrap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	router = bootstrap.Bootstrap()
	w = httptest.NewRecorder()
	m.Run()
}

func TestStore(t *testing.T) {
	request := Request.UserStoreRequest{
		Name:     "zhan",
		Password: "123456",
		Email:    "390961827@qq.com",
	}
	requestStr, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(string(requestStr)))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := Response.Parse(w.Body.Bytes())
	assert.Equal(t, Response.CodeSuccess, response.Code)
	Tool.Dump(response)
}
