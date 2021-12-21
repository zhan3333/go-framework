package lgo

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-framework/internal/model"
)

type Context struct {
	*gin.Context
	User *model.User
}

func WithContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctx", &Context{
			Context: c,
		})
		c.Next()
	}
}

// HandlerFunc 控制器需要实现的方法
type HandlerFunc func(ctx *Context)

func Route(r HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.MustGet("ctx").(*Context)
		r(ctx)
		if len(ctx.Errors) != 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ctx.Errors.JSON())
			return
		}
		return
	}
}

func (c *Context) WriteErr(err error) {
	switch err.(type) {
	case *HTTPError:
		he := err.(*HTTPError)
		c.JSON(he.Code, map[string]interface{}{
			"message": he.Message,
		})
	default:
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return
}

func (c *Context) OK() error {
	return nil
}

func (c *Context) BadRequest(message ...interface{}) error {
	he := NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadGateway))
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func (c *Context) NoAuth(message ...interface{}) error {
	he := NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusBadGateway))
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func (c *Context) Bind(obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		return NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
