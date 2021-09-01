package lgo

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-framework/internal/model"
	"net/http"
)

type Context struct {
	context.Context
	Gin  *gin.Context
	User *model.User
}

func WithContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(&Context{
			c.Request.Context(),
			c,
			nil,
		})
		c.Next()
	}
}

type HandlerFunc func(ctx *Context) error

func Route(r HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context().(*Context)
		if err := r(ctx); err != nil {
			ctx.WriteErr(err)
		}
		return
	}
}

func (c *Context) JSON(body interface{}) error {
	c.Gin.JSON(http.StatusOK, body)
	return nil
}

func (c *Context) WriteErr(err error) {
	switch err.(type) {
	case *HTTPError:
		he := err.(*HTTPError)
		c.Gin.JSON(he.Code, map[string]interface{}{
			"message": he.Message,
		})
	default:
		c.Gin.JSON(http.StatusInternalServerError, map[string]interface{}{
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
	if err := c.Gin.ShouldBind(obj); err != nil {
		return NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
