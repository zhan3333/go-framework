package ctx

import (
	"net/http"

	auth2 "go-framework/core/auth"
	"go-framework/core/lgo"
	"go-framework/internal/domain"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"go-framework/pkg/auth"
)

func Register(ctx *lgo.Context, req RegisterReq) error {
	var err error
	if isUsed, err := domain.NewUser().IsEmailUsed(req.Email); err != nil {
		return err
	} else if isUsed {
		return lgo.NewHTTPError(http.StatusBadRequest, "邮箱已被使用")
	}
	// 调用领域
	params := repo.StoreUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if _, err = domain.NewUser().Store(params); err != nil {
		// 处理错误
		return err
	}
	return ctx.OK()
}

func Login(ctx *lgo.Context, req LoginReq) error {
	var (
		err  error
		user *model.User
	)
	user, err = repo.NewUser().FirstUserByEmail(req.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return ctx.BadRequest("用户不存在")
	}
	if err = auth.Compare(user.Password, req.Password); err != nil {
		return ctx.BadRequest("密码不正确")
	}
	if token, err := auth2.NewJWT().Create(uint64(user.ID)); err != nil {
		return err
	} else {
		ctx.JSON(http.StatusOK, LoginResp{
			AccessToken: token.Token,
			Type:        token.Type,
			ExpiresAt:   token.ExpiresAt,
		})
		return nil
	}
}
