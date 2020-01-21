package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/Models"
	"go-framework/app/Services"
	"log"
)

type userController struct {
}

var UserController = new(userController)

func (UserController userController) Store(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	if Services.EmailExists(email) {
		c.JSON(200, Response{
			Code:    1,
			Message: "Email has been used",
			Body:    nil,
		})
		return
	}
	user := Models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	log.Println(user)
	Models.DB.Conn.Create(&user)
	c.JSON(200, Response{
		Code:    0,
		Message: "Success",
		Body:    user,
	})
}
