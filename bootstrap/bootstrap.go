package bootstrap

import (
	"fmt"
	"go-framework/app"
	"go-framework/db"
	"go-framework/log"
	"go-framework/routes"
	"os"
)

func SetInTest() {
	p, _ := os.Getwd()
	fmt.Printf("bootstrap path is %s \n", p)
	app.App.InTest = true
}

// 应用启动入口
func Bootstrap() {
	//migrate()

	app.Init()

	log.Init()

	db.Init()

	routes.InitRouter()
}

func Destroy() {
	defer db.Close()
}

//func migrate() {
//	db.Conn.AutoMigrate(&user.User{})
//}
