package auth

import (
	jwtgo "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go-framework/app"
	"go-framework/pkg/jwt"
	"time"
)

var (
	ErrTimeoutToken = errors.New("登录凭据过期")
	ErrTokenFormat  = errors.New("无效的凭据格式")
	ErrNoToken      = errors.New("未提供登录凭据")
	ErrUserNoExists = errors.New("用户不存在")
)

func NewJWT() JWT {
	return JWT{
		Secret:      app.Config.JWT.Secret,
		ExpDuration: app.Config.JWT.TTL,
		Issuer:      app.Config.JWT.Issuer,
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

type JWTToken struct {
	Token     string
	Type      string
	ExpiresAt int64
}

func (A JWT) Create(userID uint64) (*JWTToken, error) {
	var err error
	var expiresAt = time.Now().Add(A.ExpDuration).Unix()
	j := jwt.NewJWT(A.Secret)
	token, err := j.Create(JWTClaims{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    A.Issuer,
			ExpiresAt: expiresAt,
		},
		UserID:     userID,
		Authorized: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "创建 token 失败")
	}
	return &JWTToken{
		Token:     token,
		Type:      "bearer",
		ExpiresAt: expiresAt,
	}, nil
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
