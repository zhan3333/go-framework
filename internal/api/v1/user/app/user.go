package app

import (
	"go-framework/core/http/resp"
	"go-framework/internal/domain"
	"go-framework/internal/repo"
)

func UserList(_ UserListRequest, resp resp.Responser) {
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
