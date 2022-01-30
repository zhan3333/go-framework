package lgo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"go-framework/core/auth"
	"go-framework/internal/config"
)

// Dependencies 服务依赖
type Dependencies struct {
	Logger *logrus.Logger
	DB     *gorm.DB
	Redis  *redis.Client
	Config *config.Config
	Server *gin.Engine
	JWT    *auth.JWT
}
