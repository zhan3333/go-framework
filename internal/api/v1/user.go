package v1

import (
	"github.com/gin-gonic/gin"
	resp2 "go-framework/core/http/resp"
	"go-framework/internal/api/v1/app"
)

// @Summary 创建新用户
// @Produce  json
// @Param user body requests.UserStoreRequest true "注册信息"
// @Success 200 {object} responses.ResponseStruct
// @Router /api/users [post]
func UserStore(c *gin.Context) {
	var (
		req  app.UserStoreRequest
		resp = c.MustGet("resp").(resp2.Responser)
	)
	// 参数绑定
	resp.MustBind(&req)
	app.UserStoreApp(req, resp)
}

// 获取用户列表
func UserList(c *gin.Context) {
	var (
		req  app.UserListRequest
		resp = c.MustGet("resp").(resp2.Responser)
	)
	resp.MustBind(&req)
	app.UserListApp(req, resp)
}
