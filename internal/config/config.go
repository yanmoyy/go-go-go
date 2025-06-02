package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
	Port  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	dbURL, err := getEnv("DB_URL")
	if err != nil {
		return nil, err
	}
	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		DBURL: dbURL,
		Port:  port,
	}
	return cfg, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
