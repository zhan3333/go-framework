package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-framework/app/Models"
	"go-framework/config"
	"go-framework/routes"
	"log"
	"os"
)

func SetInTest() {
	inTest = true
}

func loadConfig() {
	envFilePath := ".env"
	if inTest {
		envFilePath = os.ExpandEnv("$GOPATH/src/go-framework/.env")
	}
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.Init()
}

func loadRoutes(router *gin.Engine) {
	routes.Load(router)
}

func initDB() {
	Models.DB.Init()
}

func Bootstrap() *gin.Engine {
	loadConfig()
	initDB()
	router := gin.Default()
	loadRoutes(router)
	return router
}

func Destroy() {
	defer Models.DB.Close()
}

var inTest = false
