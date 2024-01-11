package main

import (
	"os"
	"strconv"
)

var (
	Port   = getEnv("PORT", "1323")
	LogLvl = getEnvInt("LOG_LVL", 1)
)

func loadConfig() error {
	return nil
}

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getEnvInt(key string, def int) int {
	val := getEnv(key, "")
	if val == "" {
		return def
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return v
}
