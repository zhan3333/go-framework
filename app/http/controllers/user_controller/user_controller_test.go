package user_controller_test

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-framework/app/http/Response"
	"go-framework/app/http/controllers/user_controller/requests"
	"go-framework/bootstrap"
	"go-framework/routes"
	"go-framework/tool"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	router = routes.GetRouter()
	w = httptest.NewRecorder()
	m.Run()
}

func TestStore(t *testing.T) {
	request := requests.UserStoreRequest{}
	_ = faker.FakeData(&request)
	requestStr, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(string(requestStr)))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := Response.Parse(w.Body.Bytes())
	assert.Equal(t, Response.CodeSuccess, response.Code)
	tool.Dump(response)
}

func TestList(t *testing.T) {
	request := requests.UserListRequest{}
	_ = faker.FakeData(&request)
	requestStr, _ := json.Marshal(request)
	req, _ := http.NewRequest("GET", "api/users", strings.NewReader(string(requestStr)))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := Response.Parse(w.Body.Bytes())
	assert.Equal(t, Response.CodeSuccess, response.Code)
	tool.Dump(response)
}
