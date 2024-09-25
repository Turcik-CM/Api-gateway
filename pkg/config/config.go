package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	POST_SERVICE string
	USER_SERVICE string

	ACCESS_TOKEN  string
	REFRESH_TOKEN string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}

	config.POST_SERVICE = cast.ToString(coalesce("POST_SERVICE", ":7070"))
	config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ":50050"))
	config.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "hello world"))
	config.REFRESH_TOKEN = cast.ToString(coalesce("REFRESH_TOKEN", "dodi"))
	return config

}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
