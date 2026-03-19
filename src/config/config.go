// src/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App        AppConfig
	Merchant   ServerConfig
	Retailer   ServerConfig
	APIBaseURL string
}

type AppConfig struct {
	Env string
}

type ServerConfig struct {
	Port string
}

var cfg *Config

func Load() *Config {
	if cfg != nil {
		return cfg
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	cfg = &Config{
		App: AppConfig{
			Env: getEnv("APP_ENV", "development"),
		},
		Merchant: ServerConfig{
			Port: getEnv("MERCHANT_PORT", "4001"),
		},
		Retailer: ServerConfig{
			Port: getEnv("RETAILER_PORT", "4002"),
		},
		APIBaseURL: getEnv("API_BASE_URL", "http://localhost:8080/api/v1"),
	}

	return cfg
}

func (c *Config) IsProd() bool {
	return c.App.Env == "production"
}

func (c *Config) IsDev() bool {
	return c.App.Env == "development"
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Warning: %s not set, using default: %s", key, fallback)
	return fallback
}
