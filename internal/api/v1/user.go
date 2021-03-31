package v1

import (
	"github.com/gin-gonic/gin"
	resp2 "go-framework/core/http/resp"
	"go-framework/internal/domain"
	"go-framework/internal/repo"
)

// @Summary 创建新用户
// @Produce  json
// @Param user body requests.UserStoreRequest true "注册信息"
// @Success 200 {object} responses.ResponseStruct
// @Router /api/users [post]
func UserStore(c *gin.Context) {
	var (
		err  error
		req  UserStoreRequest
		resp = c.MustGet("resp").(resp2.Responser)
	)
	// 参数绑定
	resp.MustBind(&req)
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

// 获取用户列表
func UserList(c *gin.Context) {
	var (
		req   UserListRequest
		err   error
		users []repo.ApiUser
		resp  = c.MustGet("resp").(resp2.Responser)
	)
	resp.MustBind(&req)
	if users, err = domain.NewUser().List(); err != nil {
		resp.ErrorEmpty(err)
		return
	}
	resp.SuccessBody(users)
}
