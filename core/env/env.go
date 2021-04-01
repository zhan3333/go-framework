package env

import (
	"log"
	"os"
	"strconv"
)

func DefaultGet(key string, def interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func DefaultGetBool(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatal(err)
	}
	return boolVal
}

func DefaultGetInt(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	intVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(intVal)
}
