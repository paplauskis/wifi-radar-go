package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"wifi-radar-go/internal/constant"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Address string
}

type DatabaseConfig struct {
	DatabaseDriver           string
	DatabaseConnectionString string
}

func NewConfig() *Config {
	err := godotenv.Load("internal/config/dev.env")

	if err != nil {
		panic("Error loading .env file")
	}

	c := &Config{
		Server: ServerConfig{
			Address: GetEnvOrPanic(constant.EnvKeys.ServerAddress),
		},
		Database: DatabaseConfig{
			DatabaseDriver:           GetEnvOrPanic(constant.EnvKeys.DBDriver),
			DatabaseConnectionString: GetEnvOrPanic(constant.EnvKeys.DBConnectionString),
		},
	}

	return c
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}

	return value
}
