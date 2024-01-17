package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Port   = getEnv("PORT", "1323")
	LogLvl = getEnvInt("LOG_LVL", 1)

	DBUser        = getEnv("DB_USER", "root")
	DBPass        = getEnv("DB_PASS", "root")
	DBServer      = getEnv("DB_SERVER", "localhost")
	DBPort        = getEnv("DB_PORT", "8889")
	DBName        = getEnv("DB_NAME", "guitarcomps")
	DBTablePrefix = getEnv("DB_TABLE_PREFIX", "wp")

	SigningKey = getEnvByteSlice("SIGNING_KEY", []byte("secret"))
	AppUser    = getEnv("APP_USER", "admin")
	AppPass    = getEnv("APP_PASS", "admin")
)

func loadConfig() error {
	if DBUser == "" {
		return missingConfig("DB_USER")
	}

	if DBPass == "" {
		return missingConfig("DB_PASS")
	}

	if DBServer == "" {
		return missingConfig("DB_SERVER")
	}

	if DBPort == "" {
		return missingConfig("DB_PORT")
	}

	if DBName == "" {
		return missingConfig("DB_NAME")
	}

	if SigningKey == nil {
		return missingConfig("SIGNING_KEY")
	}

	if AppUser == "" {
		return missingConfig("APP_USER")
	}

	if AppPass == "" {
		return missingConfig("APP_PASS")
	}

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

func getEnvByteSlice(key string, def []byte) []byte {
	val := getEnv(key, "")
	if val == "" {
		return def
	}
	return []byte(val)
}

func missingConfig(key string) error {
	return fmt.Errorf("missing %s", key)
}
