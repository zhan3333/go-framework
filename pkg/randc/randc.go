package randc

import (
	uuid2 "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 产生指定长度的随机字符串
func RandStringN(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Uuid() string {
	return strings.ReplaceAll(uuid2.NewV4().String(), "-", "")
}
