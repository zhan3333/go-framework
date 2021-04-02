package user_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	app2 "go-framework/app"
	_ "go-framework/core/boot/http"
	"go-framework/core/http/resp"
	"go-framework/internal/api/v1/user/app"
	"go-framework/internal/subdomain"
	"go-framework/pkg/test"
	"go-framework/pkg/util"
	"net/http"
	"testing"
)

var httpTest = test.New(app2.GetRouter())

func TestList(t *testing.T) {
	token, err := subdomain.NewAuth().EmailToLoginToken("MgWAcuH@wRLsT.info")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	request := app.UserListRequest{}
	assert.Nil(t, faker.FakeData(&request))

	httpResp := httpTest.Get("/api/v1/users", request, test.Header{
		Key:   "Authorization",
		Value: token,
	})

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeSuccess, response.Code)
	util.Dump(response)
}
