package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-framework/core/auth"
	"go-framework/core/lgo"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var user *model.User
		var err error

		ctx := c.Request.Context().(*lgo.Context)
		if token = GetToken(c); token == "" {
			ctx.WriteErr(ctx.NoAuth("未提供登录凭据"))
			c.Abort()
			return
		}
		user, err = ParseUser(token)
		if err != nil {
			ctx.WriteErr(err)
			c.Abort()
			return
		}
		ctx.User = user
		c.Next()
	}
}

func GetToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if strings.Contains(token, "bearer ") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	return token
}

func ParseUser(token string) (*model.User, error) {
	var user *model.User
	var err error
	claims, err := auth.NewJWT().Parse(token)
	if err != nil {
		return nil, err
	}
	userID := claims.UserID
	if user, err = repo.NewUser().First(userID); err != nil {
		return nil, err
	}
	fmt.Println(user, err, userID)
	if user == nil {
		return nil, auth.ErrUserNoExists
	}
	return user, nil
}
