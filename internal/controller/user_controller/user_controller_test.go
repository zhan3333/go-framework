package user_controller_test

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-framework/boot"
	"go-framework/internal/controller/user_controller/requests"
	"go-framework/internal/responses"
	routes "go-framework/internal/route"
	"go-framework/pkg/tool"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
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
	response := responses.Parse(w.Body.Bytes())
	assert.Equal(t, responses.CodeSuccess, response.Code)
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
	response := responses.Parse(w.Body.Bytes())
	assert.Equal(t, responses.CodeSuccess, response.Code)
	tool.Dump(response)
}
