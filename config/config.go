package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppVersion string
	Env        string
	Port       string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env file: %v", err)
	}

	appName := os.Getenv("APP_NAME")
	appVersion := os.Getenv("APP_VERSION")
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	return &Config{
		AppName:    appName,
		AppVersion: appVersion,
		Env:        env,
		Port:       port,
	}, nil
}

func (c *Config) Addr() string {
	return ":" + c.Port
}
