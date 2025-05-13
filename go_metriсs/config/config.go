package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		timeDuration, err := time.ParseDuration(value)
		if err == nil {
			return timeDuration
		}
		log.Printf("Error parsing environment variable %s: %v.", key, err)
	}
	return defaultValue
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	return &Config{
		Port:         getEnv("PORT", "8080"),
		TimeDuration: getEnvAsDuration("TIME_DURATION", 500*time.Millisecond),
	}
}
