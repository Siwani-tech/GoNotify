package config

import (
	"os"
	"strconv"
)

func GetEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}

	return num
}
