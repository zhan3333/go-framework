package jwt

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

//define customize claims exp:
//  ```
//	type AuthJWTClaims struct {
//		jwt2.StandardClaims
//		UserID     uint64
//		Authorized bool
//	}
//  ```
//create token exp:
//  ```
//	j := jwt.NewJWT("123456", time.Minute*15)
//	claims := AuthJWTClaims{
//		UserID:     uint64(1),
//		Authorized: true,
//	}
//	tokenStr, err := j.Create(claims)
//  ```
//parse token exp:
// claims2 will filling parse data
//  ```
//	claims2 := AuthJWTClaims{}
//	token, err := j.Parse(tokenStr, &claims2)
//  ```

type JWTer interface {
	Create(claims jwtgo.Claims) (string, error)
	Parse(tokenStr string, claims jwtgo.Claims) (*jwtgo.Token, error)
}

func NewJWT(secret string) JWTer {
	return JWT{
		Secret: secret,
	}
}

type JWT struct {
	Secret string
}

func (J JWT) Create(claims jwtgo.Claims) (string, error) {
	var err error
	now := time.Now()
	if s, ok := claims.(jwtgo.StandardClaims); ok {
		s.IssuedAt = now.Unix()
		claims = s
	}
	if s, ok := claims.(jwtgo.MapClaims); ok {
		s["iat"] = now.Unix()
		claims = s
	}
	at := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(J.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (J JWT) Parse(tokenStr string, claims jwtgo.Claims) (*jwtgo.Token, error) {
	token, err := jwtgo.ParseWithClaims(tokenStr, claims, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(J.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
