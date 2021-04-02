package middleware

import (
	"github.com/gin-gonic/gin"
	"go-framework/core/auth"
	resp2 "go-framework/core/http/resp"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var user *model.User
		var err error
		resp := c.MustGet("resp").(resp2.Responser)
		if token = GetToken(c); token == "" {
			resp.ErrorEmpty(auth.ErrNoToken)
			c.Abort()
			return
		}
		user, err = ParseUser(token)
		if err != nil {
			resp.ErrorEmpty(err)
			c.Abort()
			return
		}
		c.Set("user", user)
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
	if user == nil {
		return nil, auth.ErrUserNoExists
	}
	return user, nil
}
