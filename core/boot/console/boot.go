package console

import "go-framework/core/boot"

// 引入即可启动控制台模式下的框架
// 不会加载路由组件

func init() {
	boot.SetInCommand()
	_ = boot.Boot()
}
