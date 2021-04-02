package app

import (
	auth2 "go-framework/core/auth"
	"go-framework/core/http/resp"
	"go-framework/internal/domain"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"go-framework/pkg/auth"
)

func Register(req RegisterReq, resp resp.Responser) {
	var err error
	if isUsed, err := domain.NewUser().IsEmailUsed(req.Email); err != nil {
		resp.ErrorEmpty(err)
		return
	} else if isUsed {
		resp.FailedMsg("邮箱已被使用")
		return
	}
	// 调用领域
	params := repo.StoreUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if _, err = domain.NewUser().Store(params); err != nil {
		// 处理错误
		resp.ErrorEmpty(err)
		return
	}
	// 响应成功
	resp.SuccessEmpty()
}

func Login(req LoginReq, resp resp.Responser) {
	var (
		err  error
		user *model.User
	)
	user, err = repo.NewUser().FirstUserByEmail(req.Email)
	if err != nil {
		resp.ErrorEmpty(err)
		return
	}
	if user == nil {
		resp.FailedMsg("用户不存在")
		return
	}
	if err = auth.Compare(user.Password, req.Password); err != nil {
		resp.FailedMsg("密码不正确")
		return
	}
	if token, err := auth2.NewJWT().Create(uint64(user.ID)); err != nil {
		resp.ErrorEmpty(err)
		return
	} else {
		resp.SuccessBody(LoginResp{
			AccessToken: token.Token,
			Type:        token.Type,
			ExpiresAt:   token.ExpiresAt,
		})
	}
}
