package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-framework/app"
	"go-framework/app/Models/User"
	"go-framework/config"
	"go-framework/database"
	"go-framework/routes"
	"io"
	"log"
	"os"
	"path/filepath"
)

func SetInTest() {
	app.Application.InTest = true
}

func loadConfig() {
	envFilePath := ".env"
	app.Application.AppPath, _ = os.Getwd()
	if app.Application.InTest {
		app.Application.AppPath = os.ExpandEnv("$GOPATH/src/go-framework")
		envFilePath = os.ExpandEnv("$GOPATH/src/go-framework/.env")
	}
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Application.StoragePath = filepath.Join(app.Application.AppPath, "storage")
	app.Application.TestPath = filepath.Join(app.Application.AppPath, "tests")
	config.Init()
}

func loadRoutes(router *gin.Engine) {
	routes.Load(router)
}

func initDB() {
	database.Init()
}

func initLog() {
	gin.DisableConsoleColor()
	f, _ := os.Create(config.App.Logging.Gin.Log.Path)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Bootstrap() *gin.Engine {
	loadConfig()
	initLog()
	initDB()
	migrate()
	router := gin.Default()
	loadRoutes(router)
	return router
}

func Destroy() {
	defer database.Close()
}

func migrate() {
	database.Conn.AutoMigrate(&User.User{})
}
