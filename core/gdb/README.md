# gdb

gorm 多连接管理

## 介绍

gorm v2 版本多连接, 使用需要预先配置 DefaultName ConnConfigs 两项

建议项目启动时即初始化所有连接:

1. 初始化默认连接 InitDef()
2. 初始化所有连接 InitAll()

需要访问默认连接 (Entry) , 可以使用以下方法

1. gdb.DB
2. gdb.Def()
3. gdb.Conn(gdb.DefaultName)

要访问 sql.DB 对象, 可以使用

1. gdb.Conn("connection name").SQLDB

## 安装

`go get -u github.com/zhan3333/gdb/v2@v2.0.0-alpha.1`

## 使用

初始化

```go
package main

import (
	"fmt"
	"go-framework/core/gdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	maxLiftTime := 30 * time.Second
	parseTime := true
	gdb.DefaultName = "default"
	gdb.ConnConfigs = map[string]gdb.MySQLConf{
		gdb.DefaultName: {
			Host:        "127.0.0.1",
			Port:        "3306",
			Username:    "root",
			Password:    "123456",
			Database:    "test",
			ParseTime:   &parseTime,
			MaxLiftTime: &maxLiftTime,
			GORMConfig: &gorm.Config{
				Logger: logger.New(
					log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
					logger.Config{
						SlowThreshold: time.Second, // 慢 SQL 阈值
						LogLevel:      logger.Info, // Log level
						Colorful:      true,        // 禁用彩色打印
					},
				),
			},
		},
	}
	_, err := gdb.InitDef()
	if err != nil {
		panic(fmt.Sprintf("connection db failed: %+v\n", err))
	}
}
```

使用

```go
package main
import (
    "github.com/zhan3333/gdb"
    "gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
    var users = []User{
		{
			Name:     "name1",
			Email:    "a@example.com",
			Password: "123456",
		}, {
			Name:     "name2",
			Email:    "b@example.com",
			Password: "123456",
		},
	}
	gdb.Def().CreateInBatches(&users, 1)
}
```