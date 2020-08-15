package conf

import (
	gin2 "github.com/gin-gonic/gin"
	"go-framework/conf/env"
	_ "go-framework/conf/env"
	"os"
	"strings"
)

var (
	GinModel string
	Name     = os.Getenv("APP_NAME")
	Url      = os.Getenv("APP_URL")
	Env      = os.Getenv("APP_ENV")
	Debug    = env.DefaultGetBool("DEBUG", false)
	Host     = os.Getenv("APP_HOST")
)

func Init() {
	if !strings.EqualFold(Env, "local") && !strings.EqualFold(Env, "production") && !strings.EqualFold(Env, "testing") {
		panic("env APP_ENV must be: local, production, testing")
	}
	switch Env {
	case "testing":
		GinModel = gin2.TestMode
	case "local":
		GinModel = gin2.DebugMode
	case "production":
		GinModel = gin2.ReleaseMode
	}
}
