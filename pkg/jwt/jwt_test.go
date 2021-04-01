package jwt_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/jwt"
	"testing"
	"time"
)

func TestJWT_Create(t *testing.T) {
	auth := jwt.NewAuthJWT("123456", time.Minute*15)
	tokenStr, err := auth.Create(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenStr)
	claims, err := auth.Parse(tokenStr)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	authClaims := claims.Claims.(*jwt.AuthJWTClaims)
	assert.NotNil(t, authClaims)
	assert.Equal(t, uint64(1), authClaims.UserID)
	assert.Equal(t, true, authClaims.Authorized)
}
