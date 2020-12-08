package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhan3333/gdb/v2"
	"go-framework/internal/controller/user_controller/requests"
	"go-framework/internal/responses"

	"go-framework/internal/model"
	"go-framework/internal/service"
)

// @Summary 创建新用户
// @Produce  json
// @Param user body requests.UserStoreRequest true "注册信息"
// @Success 200 {object} responses.Response
// @Router /api/users [post]
func Store(c *gin.Context) {
	var request requests.UserStoreRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		responses.BadReq(c, err)
		return
	}
	if service.GetUserService().EmailIsExists(request.Email) {
		responses.Failed(c, "Email has been used", nil)
		return
	}
	u := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	err = gdb.Def().Create(&u).Error
	if err != nil {
		responses.Error(c, err)
		return
	}
	responses.Success(c, "success", u)
}

func List(c *gin.Context) {
	var request requests.UserListRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		responses.BadReq(c, err)
		return
	}
	var users []model.User
	err = gdb.Def().Find(&users).Error
	if err != nil {
		responses.Error(c, err)
		return
	}
	//log.Print(users)
	responses.Success(c, "success", users)
}
