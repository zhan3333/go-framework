package auth

import (
	"go-framework/core/lgo"
	"go-framework/internal/api/v1/auth/app"
)

// Register @Summary 注册新用户
// @Produce  json
// @Param user body app.RegisterReq true "注册信息"
// @Success 200 {object} resp2.Responser
// @Router /api/auth/register [post]
func Register(ctx *lgo.Context) error {
	var (
		req app.RegisterReq
	)
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return app.Register(ctx, req)
}

// Login @Summary 登录
// @Produce  json
// @Param user body app.LoginReq true "登录"
// @Success 200 {object} resp.LoginResp
// @Router /api/auth/login [post]
func Login(ctx *lgo.Context) error {
	var (
		req app.LoginReq
	)
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return app.Login(ctx, req)
}
