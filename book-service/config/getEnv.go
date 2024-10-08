package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	BOOK_SERVICE_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	MONGO_DB_HOST string
	MONGO_DB_PORT int
	MONGO_DB_NAME string
	MONGO_DB_USER string
	MONGO_DB_PASS string

	KAFKA_HOST string
	KAFKA_PORT string

	REDIS_HOST string
	REDIS_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.BOOK_SERVICE_PORT = cast.ToString(coalesce("BOOK_SERVICE_PORT", ":50024"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "bookshop"))

	config.MONGO_DB_HOST = cast.ToString(coalesce("MONGO_DB_HOST", "mongo"))
	config.MONGO_DB_USER = cast.ToString(coalesce("MONGO_DB_USER", "sardorbek"))
	config.MONGO_DB_PASS = cast.ToString(coalesce("MONGO_DB_PASS", "1111"))
	config.MONGO_DB_PORT = cast.ToInt(coalesce("MONGO_DB_PORT", 27017))
	config.MONGO_DB_NAME = cast.ToString(coalesce("MONGO_DB_NAME", "bookshop"))

	config.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
	config.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", ":9092"))

	config.REDIS_HOST = cast.ToString(coalesce("REDIS_HOST", "localhost"))
	config.REDIS_PORT = cast.ToString(coalesce("REDIS_PORT", ":6379"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
