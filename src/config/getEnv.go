package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BOT_TOKEN   string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using default values...")
	}

	config := Config{}

	config.BOT_TOKEN = "6652950296:AAEKXUjZbtDs4Cu_XE40tPiECMOU_Zu5iTI"

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "polling-cont"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "polling_db"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
