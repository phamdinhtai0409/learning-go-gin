package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	App struct {
		Env   string
		Debug string
		Port  int
	}

	DB struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
	}

	Jwt struct {
		SecretKey       string
		Issuer          string
		ExpirationHours int
	}
}

func NewConfig() Config {
	c := Config{}

	c.App.Env = Get("APP_ENV", "development")
	c.App.Debug = Get("APP_DEBUG", "")
	c.App.Port = GetInt("APP_PORT", 5050)

	c.DB.Host = Get("DB_HOST", "localhost")
	c.DB.Port = GetInt("DB_PORT", 5432)
	c.DB.Database = Get("DB_DATABASE", "")
	c.DB.Username = Get("DB_USERNAME", "")
	c.DB.Password = Get("DB_PASSWORD", "")

	c.Jwt.SecretKey = Get("JWT_SECRET_KEY", "")
	c.Jwt.Issuer = Get("JWT_ISSUER", "")
	c.Jwt.ExpirationHours = GetInt("JWT_EXPIRATION_HOURS", 24)

	return c
}

func Get(key string, defaultValue string) string {
	s := os.Getenv(key)

	if s == "" {
		return defaultValue
	}

	return strings.Trim(s, "\" ")
}

func GetInt(key string, defaultValue int) int {
	a := os.Getenv(key)

	if a == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(a)

	if err != nil {
		return defaultValue
	}

	return i
}
