package auth_test

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"go-framework/app"
	"go-framework/core/boot"
	"go-framework/core/gdb"
	"go-framework/core/lgo"
	"go-framework/internal/api/v1/auth/ctx"
	"go-framework/internal/domain"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var test *lgo.Test

func TestMain(m *testing.M) {
	if err := boot.New(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
		boot.WithRoutePrint(false),
	); err != nil {
		panic(err)
	}
	test = lgo.New(app.GetRouter())
	m.Run()
}

// 测试正常注册
func TestRegister(t *testing.T) {
	request := ctx.RegisterReq{}
	assert.Nil(t, faker.FakeData(&request))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	resp := test.Send(req)

	assert.Equal(t, http.StatusOK, resp.Code)

	t.Log(resp.Body.String())
}

// 测试参数校验不通过
func TestStoreParamsErr(t *testing.T) {
	request := ctx.RegisterReq{}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	resp := test.Send(req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	t.Log(resp.Body.String())
}

// 测试邮箱已被使用
func TestRegisterEmailUsed(t *testing.T) {
	user := model.User{}
	assert.Nil(t, gdb.Def().First(&user).Error)

	request := ctx.RegisterReq{
		Email:    user.Email,
		Name:     faker.Name(),
		Password: faker.Password(),
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	resp := test.Send(req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	t.Log(resp.Body.String())

	body := struct {
		Message string
	}{}
	assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
	assert.Equal(t, "邮箱已被使用", body.Message)
}

func TestLogin(t *testing.T) {
	pwd := faker.Password()
	user, err := domain.NewUser().Store(repo.StoreUserParams{
		Name:     faker.Username(),
		Email:    faker.Email(),
		Password: pwd,
	})
	assert.Nil(t, err)

	request := ctx.LoginReq{
		Email:    user.Email,
		Password: pwd,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	resp := test.Send(req)

	assert.Equal(t, http.StatusOK, resp.Code)
	t.Log(resp.Body.String())

	body := struct {
		AccessToken string `json:"access_token"`
		Type        string `json:"type"`
		ExpiresAt   int64  `json:"expires_at"`
	}{}
	assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &body))
	assert.NotEmpty(t, body.AccessToken)
	assert.NotEmpty(t, body.Type)
	assert.NotEmpty(t, body.ExpiresAt)
}
