package swag

import "github.com/gin-gonic/gin"
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files
import _ "go-framework/docs"

func LoadSwag(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
