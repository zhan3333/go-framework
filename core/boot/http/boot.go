package http

import "go-framework/core/boot"

// 引入即可启动完整框架, 会加载路由组件

func init() {
	_ = boot.Boot()
}
