package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	POST_SERVICE     string
	USER_SERVICE     string
	NATIONAL_SERVICE string
	POST_HOST        string
	USER_HOST        string
	NATIONAL_HOST    string
	API_GATEWAY      string

	ACCES_TOKEN    string
	REFRESH_TOKEN  string
	ADMIN_PASSWORD string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}

	config.API_GATEWAY = cast.ToString(coalesce("API_GATEWAY", ":8087"))
	config.POST_SERVICE = cast.ToString(coalesce("POST_SERVICE", ":50055"))
	config.POST_HOST = cast.ToString(coalesce("POST_HOST", "localhost"))
	config.USER_SERVICE = cast.ToString(coalesce("USER_SERVICE", ":50050"))
	config.USER_HOST = cast.ToString(coalesce("USER_HOST", "localhost"))
	config.NATIONAL_SERVICE = cast.ToString(coalesce("NATIONAL_SERVICE", ":7080"))
	config.NATIONAL_HOST = cast.ToString(coalesce("NATIONAL_HOST", "18.196.33.86"))
	config.ACCES_TOKEN = cast.ToString(coalesce("ACCES_TOKEN", "hello world"))
	config.REFRESH_TOKEN = cast.ToString(coalesce("REFRESH_TOKEN", "dodi"))
	config.ADMIN_PASSWORD = cast.ToString(coalesce("ADMIN_PASSWORD", "123321"))
	return config

}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
