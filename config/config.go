package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var AdminUser string
var AdminPass string
var Port string
var DbName string
var Secret string

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
}
