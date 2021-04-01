package test

import "go-framework/core/boot"

// 引入即可启动完整框架, 不会加载路由组件

func init() {
	boot.SetInTest()
	boot.SetInCommand()
	boot.Boot()
}
