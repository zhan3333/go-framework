package subdomain

import (
	"github.com/pkg/errors"
	"go-framework/core/auth"
	"go-framework/internal/repo"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (Auth) EmailToLoginToken(email string) (string, error) {
	user, err := repo.NewUser().FirstUserByEmail(email)
	if err != nil {
		return "", errors.Wrap(err, "查询用户失败")
	}
	if user == nil {
		return "", errors.Errorf("%s 用户不存在", email)
	}
	token, err := auth.NewJWT().Create(uint64(user.ID))
	if err != nil {
		return "", errors.Wrap(err, "创建 token 失败")
	}
	return token.Token, nil
}
