package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	CORS     CORSConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// Continue even if .env file is not found
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     getEnvOrDefault("DB_NAME", "kids_shop"),
		},
		CORS: CORSConfig{
			AllowedOrigins:   strings.Split(getEnvOrDefault("CORS_ALLOWED_ORIGINS", "*"), ","),
			AllowedMethods:   strings.Split(getEnvOrDefault("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"), ","),
			AllowedHeaders:   strings.Split(getEnvOrDefault("CORS_ALLOWED_HEADERS", "Content-Type,Authorization"), ","),
			AllowCredentials: getEnvOrDefault("CORS_ALLOW_CREDENTIALS", "true") == "true",
		},
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 