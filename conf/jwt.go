package conf

import (
	"go-framework/core/env"
	"os"
	"time"
)

type JWTConf struct {
	TTL    time.Duration
	Issuer string
}

var (
	JWT = JWTConf{
		TTL:    time.Duration(env.DefaultGetInt("JWT_TTL", 86400)) * time.Second,
		Issuer: os.Getenv("APP_NAME"),
	}
)
