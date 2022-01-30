package middleware

import (
	"github.com/gin-gonic/gin"

	"go-framework/pkg/lgo"

	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token string

		cc := c.MustGet(lgo.CustomContextKey).(*lgo.CustomContext)
		if token = getToken(cc); token == "" {
			_ = cc.Unauthorized("未提供登录凭据")
			c.Abort()
			return
		}
		userID, err := parseToken(cc, token)
		if err != nil {
			_ = cc.Unauthorized("无效的登陆凭据")
			c.Abort()
			return
		}
		cc.UserID = userID
		c.Next()
	}
}

func getToken(cc *lgo.CustomContext) string {
	token := cc.GetHeader("Authorization")
	if strings.Contains(token, "bearer ") {
		token = strings.ReplaceAll(token, "bearer ", "")
	}
	return token
}

func parseToken(cc *lgo.CustomContext, token string) (uint64, error) {
	var err error
	claims, err := cc.JWT.Parse(token)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
