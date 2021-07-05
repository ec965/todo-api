package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var AdminUser string
var AdminPass string
var Port string
var DbName string
var Secret string
var TokenDuration int64
var TokenIssuer string

func envFallback(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	AdminUser = envFallback("ADMIN_USER", "admin")
	AdminPass = envFallback("ADMIN_PASS", "test123")
	Port = envFallback("PORT", "8080")
	DbName = envFallback("DB_NAME", "todo.db")
	Secret = envFallback("SECRET", "very-secret")
	TokenDuration, err = strconv.ParseInt(envFallback("TOKEN_DURATION", "1200"), 10, 64)
	if err != nil {
		TokenDuration = 1200
	}
	TokenIssuer = envFallback("TOKEN_ISSUER", "api-server")
}
