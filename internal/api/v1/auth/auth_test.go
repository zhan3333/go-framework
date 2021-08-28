package auth_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	app2 "go-framework/app"
	_ "go-framework/core/boot/http"
	"go-framework/core/http/resp"
	"go-framework/internal/api/v1/auth/app"
	"go-framework/internal/domain"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"go-framework/pkg/gdb"
	"go-framework/pkg/test"
	"go-framework/pkg/util"
	"net/http"
	"testing"
)

var httpTest = test.New(app2.GetRouter())

// 测试正常注册
func TestRegister(t *testing.T) {
	request := app.RegisterReq{}
	assert.Nil(t, faker.FakeData(&request))

	httpResp := httpTest.Post("/api/v1/auth/register", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeSuccess, response.Code)
	util.Dump(response)
}

// 测试参数校验不通过
func TestStoreParamsErr(t *testing.T) {
	request := app.RegisterReq{}

	httpResp := httpTest.Post("/api/v1/auth/register", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeFailed, response.Code)
	assert.Contains(t, response.Message, "为必填字段")
	util.Dump(response)
}

// 测试邮箱已被使用
func TestRegisterEmailUsed(t *testing.T) {
	user := model.User{}
	assert.Nil(t, gdb.Def().First(&user).Error)

	request := app.RegisterReq{
		Email:    user.Email,
		Name:     faker.Name(),
		Password: faker.Password(),
	}

	httpResp := httpTest.Post("/api/v1/auth/register", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeFailed, response.Code)
	assert.Contains(t, response.Message, "邮箱已被使用")
	util.Dump(response)
}

func TestLogin(t *testing.T) {
	pwd := faker.Password()
	user, err := domain.NewUser().Store(repo.StoreUserParams{
		Name:     faker.Username(),
		Email:    faker.Email(),
		Password: pwd,
	})
	assert.Nil(t, err)

	request := app.LoginReq{
		Email:    user.Email,
		Password: pwd,
	}

	httpResp := httpTest.Post("/api/v1/auth/login", request)

	assert.Equal(t, http.StatusOK, httpResp.Code)
	response := resp.Parse(httpResp.Body.Bytes())
	assert.Equal(t, resp.CodeSuccess, response.Code)
	assert.Contains(t, response.Message, "success")
	busBody := app.LoginResp{}
	assert.Nil(t, util.MapToStruct(response.Body.(map[string]interface{}), &busBody))
	assert.NotEmpty(t, busBody.AccessToken)
	assert.NotEmpty(t, busBody.Type)
	assert.NotEmpty(t, busBody.ExpiresAt)
}
