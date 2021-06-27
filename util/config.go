package util

import (
	"log"
	"os"
)

var AppConfig Config

type Config struct {
	SentryDSN  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func init() {
	AppConfig = Config{
		SentryDSN:  os.Getenv("SENTRY_DSN"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	log.Println("Initialized configs")
}
