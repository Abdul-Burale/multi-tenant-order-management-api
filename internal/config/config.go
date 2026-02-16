package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DBConn          string
	JWTSecret       string
	StripeSecretKey string
	AppEnv          string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:            getEnv("PORT", "8080"),
		DBConn:          os.Getenv("DB_CONN"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		AppEnv:          getEnv("APP_ENV", "development"),
	}

	if cfg.DBConn == "" {
		return nil, errors.New("DB_CONN is required")
	}

	if cfg.JWTSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	if cfg.StripeSecretKey == "" {
		return nil, errors.New("STRIPE_SECRET_KEY is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
