package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/Http/Request"
	"go-framework/app/Http/Response"
	"go-framework/app/Models/User"
	"go-framework/database"
)

type userController struct {
}

var UserController = new(userController)

func (UserController userController) Store(c *gin.Context) {
	var request Request.UserStoreRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		Response.Error(c, err)
		return
	}
	if User.EmailIsExists(request.Email) {
		Response.Failed(c, "Email has been used", nil)
		return
	}
	user := User.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	err = database.Conn.Create(&user).Error
	if err != nil {
		Response.Error(c, err)
		return
	}
	Response.Success(c, "success", user)
}
