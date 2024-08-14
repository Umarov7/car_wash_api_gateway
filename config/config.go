package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT                        string
	AUTH_SERVICE_PORT                string
	BOOKING_SERVICE_PORT             string
	DB_HOST                          string
	DB_PORT                          int
	DB_USER                          string
	DB_PASSWORD                      string
	DB_NAME                          string
	ACCESS_TOKEN                     string
	KAFKA_HOST                       string
	KAFKA_PORT                       string
	KAFKA_TOPIC_BOOKING_CREATED      string
	KAFKA_TOPIC_BOOKING_UPDATED      string
	KAFKA_TOPIC_BOOKING_CANCELLED    string
	KAFKA_TOPIC_PAYMENT_CREATED      string
	KAFKA_TOPIC_REVIEW_CREATED       string
	KAFKA_TOPIC_NOTIFICATION_CREATED string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", "api-gateway:8080"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", "8081"))
	cfg.BOOKING_SERVICE_PORT = cast.ToString(coalesce("BOOKING_SERVICE_PORT", "8082"))

	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres"))
	cfg.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	cfg.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	cfg.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "password"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "postgres"))

	cfg.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "ACCESS_TOKEN"))

	cfg.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "kafka"))
	cfg.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", "9092"))

	cfg.KAFKA_TOPIC_BOOKING_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_CREATED", "car-wash.booking_created"))
	cfg.KAFKA_TOPIC_BOOKING_UPDATED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_UPDATED", "car-wash.booking_updated"))
	cfg.KAFKA_TOPIC_BOOKING_CANCELLED = cast.ToString(coalesce("KAFKA_TOPIC_BOOKING_CANCELLED", "car-wash.booking_cancelled"))
	cfg.KAFKA_TOPIC_PAYMENT_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_PAYMENT_CREATED", "car-wash.payment_created"))
	cfg.KAFKA_TOPIC_REVIEW_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_REVIEW_CREATED", "car-wash.review_created"))
	cfg.KAFKA_TOPIC_NOTIFICATION_CREATED = cast.ToString(coalesce("KAFKA_TOPIC_NOTIFICATION_CREATED", "car-wash.notification_created"))

	return cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
