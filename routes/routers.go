package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-framework/conf"
	"go-framework/routes/api"
	"go-framework/storage"
	"io"
	"os"
)

var router *gin.Engine

func InitRouter() {
	gin.SetMode(conf.Conf.GinModel)
	gin.DisableConsoleColor()
	fmt.Printf("%s \n", storage.Storage.AbsPath)
	f, _ := os.Create(storage.Storage.AbsPath)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router = gin.Default()
	router.Use(gin.Recovery())
	api.LoadApi(router)
}

func GetRouter() *gin.Engine {
	if router == nil {
		InitRouter()
	}
	return router
}
