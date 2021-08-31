package user

import (
	"go-framework/core/lgo"
	"go-framework/internal/api/v1/user/app"
)

// List 获取用户列表
func List(ctx *lgo.Context) error {
	var (
		req app.UserListRequest
	)
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return app.UserList(ctx, req)
}
