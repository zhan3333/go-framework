package user

import (
	"github.com/gin-gonic/gin"
	resp2 "go-framework/core/http/resp"
	"go-framework/internal/api/v1/user/app"
)

// 获取用户列表
func List(c *gin.Context) {
	var (
		req  app.UserListRequest
		resp = c.MustGet("resp").(resp2.Responser)
	)
	resp.MustBind(&req)
	app.UserList(req, resp)
}
