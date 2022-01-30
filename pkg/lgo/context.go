package lgo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	"go-framework/internal/domain/user"
)

const CustomContextKey = "custom_context"

type CustomContext struct {
	*gin.Context
	UserID uint64
	*Dependencies
}

func WithContext(dependencies *Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(CustomContextKey, &CustomContext{
			Context:      c,
			Dependencies: dependencies,
		})
		c.Next()
	}
}

// CustomHandlerFunc 控制器需要实现的方法
type CustomHandlerFunc func(ctx *CustomContext) error

func Controller(r CustomHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet(CustomContextKey).(*CustomContext)
		if err := r(ctx); err != nil {
			ctx.WriteErr(err)
		} else {
			if ctx.Writer.Status() == 0 {
				_ = ctx.OK()
			}
		}
	}
}

// WriteErr 向响应中写入错误
func (c *CustomContext) WriteErr(err error) {
	var he *HTTPError
	if errors.As(err, &he) {
		_ = c.Message(he.Code, he.Message)
	} else {
		_ = c.Message(http.StatusInternalServerError, err.Error())
	}
}

func (c *CustomContext) Message(code int, message ...string) error {
	msg := http.StatusText(code)
	if len(message) > 0 {
		msg = message[0]
	}
	c.JSON(code, gin.H{
		"message": msg,
	})
	return nil
}

func (c *CustomContext) OK(v ...interface{}) error {
	if len(v) > 0 {
		c.JSON(http.StatusOK, v[0])
		return nil
	}
	return c.Message(http.StatusOK, "ok")
}

func (c *CustomContext) BadRequest(message ...string) error {
	return c.Message(http.StatusBadRequest, message...)
}

func (c *CustomContext) Unauthorized(message ...string) error {
	return c.Message(http.StatusUnauthorized, message...)
}

func (c *CustomContext) InternalServerError(message ...string) error {
	return c.Message(http.StatusInternalServerError, message...)
}

func (c *CustomContext) Bind(obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		return c.BadRequest(err.Error())
	}
	return nil
}

func (c *CustomContext) NewUser() *user.User {
	return user.NewUser(c.DB)
}
