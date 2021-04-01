package conf

import (
	"github.com/zhan3333/gdb/v2"
	"github.com/zhan3333/gredis"
	"go-framework/core/env"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type database struct {
	MySQL map[string]gdb.MySQLConf
	Redis map[string]gredis.Conf
}

var Database = database{
	MySQL: map[string]gdb.MySQLConf{
		gdb.DefaultName: {
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			Username:   os.Getenv("DB_USERNAME"),
			Password:   os.Getenv("DB_PASSWORD"),
			Database:   os.Getenv("DB_DATABASE"),
			GORMConfig: &gorm.Config{},
		},
	},
	Redis: map[string]gredis.Conf{
		gredis.DefaultConn: {
			Host:     env.DefaultGet("REDIS_HOST", "127.0.0.1").(string),
			Password: env.DefaultGet("REDIS_PASSWORD", "").(string),
			Port:     env.DefaultGetInt("REDIS_PORT", 6379),
			Database: env.DefaultGetInt("REDIS_DATABASE", 0),
		},
	},
}

func dbLogMode() logger.LogLevel {
	if env.DefaultGetBool("DB_LOG_MODE", true) {
		return logger.Info
	}
	return logger.Silent
}

//由于 glog 包加载顺序与本包不确定, 故需要手动设置日志类
func InitDBLogMode() {
	mode := dbLogMode()
	Database.MySQL[gdb.DefaultName].GORMConfig.Logger = logger.New(
		// 使用 glog 记录, 需要关闭日志配置的 format
		// glog.Def(),
		// 输出到屏幕, 可以开启颜色选项
		log.New(os.Stdout, "", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      mode,        // Log level
			Colorful:      true,        // 彩色打印
		},
	)
}
