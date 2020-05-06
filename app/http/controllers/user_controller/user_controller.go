package user_controller

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/http/Response"
	"go-framework/app/http/controllers/user_controller/requests"
	"go-framework/db"
	"go-framework/models/user"
)

type userController struct {
}

var UserController = new(userController)

func (userController userController) Store(c *gin.Context) {
	var request requests.UserStoreRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		Response.Error(c, err)
		return
	}
	if user.EmailIsExists(request.Email) {
		Response.Failed(c, "Email has been used", nil)
		return
	}
	u := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	err = db.Conn.Create(&u).Error
	if err != nil {
		Response.Error(c, err)
		return
	}
	Response.Success(c, "success", u)
}

func (userController userController) List(c *gin.Context) {
	var request requests.UserListRequest
	var err error
	if err = c.ShouldBindJSON(&request); err != nil {
		Response.Error(c, err)
		return
	}
	var users []user.User
	err = db.Conn.Find(&users).Error
	if err != nil {
		Response.Error(c, err)
		return
	}
	//log.Print(users)
	Response.Success(c, "success", users)
}
