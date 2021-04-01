package env

import (
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	// 嵌套查找父文件夹中的 .env 文件
	envFilePath := ".env"
	errCount := 0
	success := false
	for errCount < 10 {
		err := godotenv.Load(envFilePath)
		if err != nil {
			envFilePath = fmt.Sprintf("../%s", envFilePath)
		} else {
			success = true
			break
		}
		errCount++
	}
	if !success {
		panic("Error loading .env file")
	}
}
