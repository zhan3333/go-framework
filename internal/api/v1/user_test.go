package v1_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go-framework/app"
	_ "go-framework/core/boot/http"
	"go-framework/core/http/resp"
	app2 "go-framework/internal/api/v1/app"
	"go-framework/pkg/test"
	"go-framework/pkg/tool"
	"net/http"
	"testing"
)

var httpTest = test.New(app.GetRouter())

func TestStore(t *testing.T) {
	request := app2.UserStoreRequest{}
	assert.Nil(t, faker.FakeData(&request))

	httpResp := httpTest.Post("/api/v1/users", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeSuccess, response.Code)
	tool.Dump(response)
}

func TestStoreParamsErr(t *testing.T) {
	request := app2.UserStoreRequest{}

	httpResp := httpTest.Post("/api/v1/users", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeFailed, response.Code)
	assert.Contains(t, response.Message, "为必填字段")
	tool.Dump(response)
}

func TestList(t *testing.T) {
	request := app2.UserListRequest{}
	assert.Nil(t, faker.FakeData(&request))

	httpResp := httpTest.Get("/api/v1/users", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeSuccess, response.Code)
	tool.Dump(response)
}
