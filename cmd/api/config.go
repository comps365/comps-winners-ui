package main

import (
	"errors"
	"os"
	"strconv"
)

var (
	Port   = getEnv("PORT", "1323")
	LogLvl = getEnvInt("LOG_LVL", 1)

	DBUser        = getEnv("DB_USER", "")
	DBPass        = getEnv("DB_PASS", "")
	DBServer      = getEnv("DB_SERVER", "")
	DBPort        = getEnv("DB_PORT", "")
	DBName        = getEnv("DB_NAME", "")
	DBTablePrefix = getEnv("DB_TABLE_PREFIX", "wp")

	dbUserRequired   = errors.New("DB_USER is required")
	dbPassRequired   = errors.New("DB_PASS is required")
	dbServerRequired = errors.New("DB_SERVER is required")
	dbPortRequired   = errors.New("DB_PORT is required")
	dbNameRequired   = errors.New("DB_NAME is required")
)

func loadConfig() error {
	if DBUser == "" {
		return dbUserRequired
	}

	if DBPass == "" {
		return dbPassRequired
	}

	if DBServer == "" {
		return dbServerRequired
	}

	if DBPort == "" {
		return dbPortRequired
	}

	if DBName == "" {
		return dbNameRequired
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
