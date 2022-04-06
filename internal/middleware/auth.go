package middleware

import (
	"github.com/gin-gonic/gin"

	"go-framework/pkg/lgo"

	"strings"
)

const (
	NoToken      = "未提供登录凭据"
	InvalidToken = "无效的登陆凭据"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token string

		cc := c.MustGet(lgo.CustomContextKey).(*lgo.CustomContext)
		if token = getToken(cc); token == "" {
			_ = cc.Unauthorized(NoToken)
			c.Abort()
			return
		}
		userID, err := parseToken(cc, token)
		if err != nil {
			_ = cc.Unauthorized(InvalidToken + ": " + err.Error())
			c.Abort()
			return
		}
		cc.UserID = userID
		c.Next()
	}
}

func getToken(c *lgo.CustomContext) string {
	token := c.GetHeader("Authorization")
	if strings.Contains(token, "bearer ") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	return token
}

func parseToken(c *lgo.CustomContext, token string) (uint64, error) {
	var err error
	claims, err := c.JWT.Parse(token)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
