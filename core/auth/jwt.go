package auth

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"go-framework/conf"
	"go-framework/pkg/jwt"
	"time"
)

var (
	ErrTimeoutToken = errors.New("登录凭据过期")
	ErrTokenFormat  = errors.New("无效的凭据格式")
)

func NewJWT() JWT {
	return JWT{
		Secret:      conf.Secret,
		ExpDuration: conf.JWT.TTL,
		Issuer:      conf.JWT.Issuer,
	}
}

type JWT struct {
	Secret      string
	ExpDuration time.Duration
	Issuer      string
}

type JWTClaims struct {
	jwtgo.StandardClaims
	UserID     uint64
	Authorized bool
}

func (A JWT) Create(userID uint64) (string, error) {
	var err error
	now := time.Now()
	j := jwt.NewJWT(A.Secret)
	token, err := j.Create(JWTClaims{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    A.Issuer,
			ExpiresAt: now.Add(A.ExpDuration).Unix(),
		},
		UserID:     userID,
		Authorized: true,
	})
	if err != nil {
		return "", errors.Wrap(err, "创建 token 失败")
	}
	return token, nil
}

func (A JWT) Parse(tokenStr string) (JWTClaims, error) {
	claims := JWTClaims{}
	j := jwt.NewJWT(A.Secret)
	token, err := j.Parse(tokenStr, &claims)
	if err != nil {
		return claims, err
	}
	if err = token.Claims.Valid(); err != nil {
		e := err.(jwtgo.ValidationError)
		switch e.Errors {
		case jwtgo.ValidationErrorExpired:
			return claims, ErrTimeoutToken
		case jwtgo.ValidationErrorMalformed:
			return claims, ErrTokenFormat
		default:
			return claims, errors.New(e.Error())
		}
	}
	return claims, nil
}
