package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/Models/User"
	"go-framework/database"
)

type homeController struct {
}

var HomeController = new(homeController)

func (homeController) Index(c *gin.Context) {
	users := []User.User{}
	database.Conn.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}
