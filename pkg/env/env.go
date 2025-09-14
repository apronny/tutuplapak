package env

import (
	"os"
	"strconv"
)

func String(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func Int(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if val, err := strconv.Atoi(val); err == nil {
			return val
		}
	}
	return fallback
}

func Bool(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if val, err := strconv.ParseBool(val); err == nil {
			return val
		}
	}
	return fallback
}
