package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-framework/app/Models"
)

type homeController struct {
}

var HomeController = new(homeController)

func (homeController) Index(c *gin.Context) {
	users := []Models.User{}
	Models.DB.Conn.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}
