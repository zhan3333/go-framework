package user

import (
	"github.com/gin-gonic/gin"
	resp2 "go-framework/core/http/resp"
	"go-framework/internal/api/v1/user/app"
)

// @Summary 注册新用户
// @Produce  json
// @Param user body app.RegisterReq true "注册信息"
// @Success 200 {object} resp2.Responser
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var (
		req  app.RegisterReq
		resp = c.MustGet("resp").(resp2.Responser)
	)
	// 参数绑定
	resp.MustBind(&req)
	app.Register(req, resp)
}

// 获取用户列表
func List(c *gin.Context) {
	var (
		req  app.UserListRequest
		resp = c.MustGet("resp").(resp2.Responser)
	)
	resp.MustBind(&req)
	app.UserList(req, resp)
}
