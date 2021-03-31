package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson/buffer"
	resp2 "go-framework/core/http/resp"
	validator2 "go-framework/internal/validator"
	"gopkg.in/go-playground/validator.v9"
)

// 在上下文中加入 resp 对象, 提供便捷的响应数据方式
// 当调用 resp.MustBind() 绑定参数失败时, 将由中间件来处理绑定失败响应

func WithResp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp = resp2.NewResp(c)
		defer func() {
			// 处理绑定错误
			if err := recover(); err != nil {
				if validErr, ok := err.(validator.ValidationErrors); ok {
					var errMsg buffer.Buffer
					errData := validErr.Translate(validator2.Trans)
					// 取第一个错误信息
					for _, value := range errData {
						errMsg.AppendString(value)
						break
					}
					resp.FailedMsg(string(errMsg.Buf))
					c.Abort()
					return
				}
				panic(err)
			}
		}()
		c.Set("resp", resp)
		c.Next()
	}
}
