package app

import (
	"go-framework/core/http/resp"
	"go-framework/internal/domain"
	"go-framework/internal/repo"
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

func UserList(req UserListRequest, resp resp.Responser) {
	var (
		err   error
		users []repo.ApiUser
	)
	if users, err = domain.NewUser().List(); err != nil {
		resp.ErrorEmpty(err)
		return
	}
	resp.SuccessBody(users)
}
