package auth_test

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"go-framework/internal/api/v1/auth"
	"go-framework/internal/domain/user"
	"go-framework/internal/model"
	"go-framework/pkg/boot"
	"go-framework/pkg/lgo"
	"go-framework/pkg/test"

	"net/http"
	"os"
	"testing"
)

var booted *boot.Booted
var server *gin.Engine

func TestMain(m *testing.M) {
	var err error
	if booted, err = boot.Boot(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
	); err != nil {
		panic(err)
	}
	server = booted.Server
	m.Run()
}

// 测试正常注册
func TestRegister(t *testing.T) {
	request := auth.RegisterReq{}
	assert.Nil(t, faker.FakeData(&request))

	req, rec := test.NewHTTPTest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	server.ServeHTTP(rec, req)
	test.PrintRec(t, rec)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// 测试参数校验不通过
func TestStoreParamsErr(t *testing.T) {
	request := auth.RegisterReq{}

	req, rec := test.NewHTTPTest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	server.ServeHTTP(rec, req)
	test.PrintRec(t, rec)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// 测试邮箱已被使用
func TestRegisterEmailUsed(t *testing.T) {
	u := model.User{}
	assert.Nil(t, booted.DB.First(&u).Error)

	request := auth.RegisterReq{
		Email:    u.Email,
		Name:     faker.Name(),
		Password: faker.Password(),
	}

	req, rec := test.NewHTTPTest(http.MethodPost, "/api/v1/auth/register", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	server.ServeHTTP(rec, req)
	test.PrintRec(t, rec)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	body := struct {
		Message string
	}{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
	assert.Equal(t, "邮箱已被使用", body.Message)
}

func NewUser() *user.User {
	return user.NewUser(booted.DB)
}

func TestLogin(t *testing.T) {
	pwd := faker.Password()
	u, err := NewUser().Store(user.StoreUserParams{
		Name:     faker.Username(),
		Email:    faker.Email(),
		Password: pwd,
	})
	assert.Nil(t, err)

	request := auth.LoginReq{
		Email:    u.Email,
		Password: pwd,
	}

	req, rec := test.NewHTTPTest(http.MethodPost, "/api/v1/auth/login", lgo.ToReader(request))
	req.Header.Set("content-type", "application/json")
	server.ServeHTTP(rec, req)

	test.PrintRec(t, rec)
	assert.Equal(t, http.StatusOK, rec.Code)

	body := struct {
		AccessToken string `json:"access_token"`
		Type        string `json:"type"`
		ExpiresAt   int64  `json:"expires_at"`
	}{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
	assert.NotEmpty(t, body.AccessToken)
	assert.NotEmpty(t, body.Type)
	assert.NotEmpty(t, body.ExpiresAt)
}

func TestMe(t *testing.T) {
	var u1 model.User
	err := booted.DB.Find(&u1).Error
	if assert.NoError(t, err) {
		token, err := booted.JWT.Create(uint64(u1.ID))
		if assert.NoError(t, err) {
			req, rec := test.NewHTTPTest(http.MethodGet, "/api/v1/me", nil)
			req.Header.Set("Authorization", token.Type+" "+token.Token)
			server.ServeHTTP(rec, req)
			test.PrintRec(t, rec)
			if assert.Equal(t, http.StatusOK, rec.Code) {
				u2 := model.User{}
				if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &u2)) {
					assert.Equal(t, u1.ID, u2.ID)
					assert.Equal(t, u1.Name, u2.Name)
					assert.Equal(t, u1.Email, u2.Email)
				}
			}
		}
	}
}
