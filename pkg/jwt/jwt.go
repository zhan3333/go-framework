package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthJWTer interface {
	Create(userID uint64) (string, error)
	Parse(token string) (*jwt.Token, error)
}

func NewAuthJWT(secret string, duration time.Duration) AuthJWT {
	return AuthJWT{
		Secret:      secret,
		ExpDuration: duration,
	}
}

type AuthJWT struct {
	AuthJWTer
	Secret      string
	ExpDuration time.Duration
}

type AuthJWTClaims struct {
	jwt.StandardClaims
	UserID     uint64
	Authorized bool
}

func (A AuthJWT) Create(userID uint64) (string, error) {
	var err error
	atClaims := AuthJWTClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(A.ExpDuration).Unix(),
		},
		userID,
		true,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(A.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (A AuthJWT) Parse(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(A.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
