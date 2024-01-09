package main

import (
	"os"
)

var (
	Port = getEnv("PORT", "1323")
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
