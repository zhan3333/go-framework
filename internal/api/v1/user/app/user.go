package app

import (
	"go-framework/core/lgo"
	"go-framework/internal/domain"
	"go-framework/internal/repo"
)

func UserList(ctx *lgo.Context, _ UserListRequest) error {
	var (
		err   error
		users []repo.ApiUser
	)
	if users, err = domain.NewUser().List(); err != nil {
		return err
	}
	return ctx.JSON(users)
}
