package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseSslMode  string
	DatabaseName     string
	Sslrootcert      string
	ServerPort       string
	JWTSecret        string
	AllowedOrigins   string
}

var CFG *Config

func LoadConfig() {
	loadEnv()
	confLogger()
}

func loadEnv() {
	if os.Getenv("GO_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found")
		}
	}

	CFG = &Config{
		DatabaseUser:     getEnv("DATABASE_USER"),
		DatabasePassword: getEnv("DATABASE_PASSWORD"),
		DatabaseHost:     getEnv("DATABASE_HOST"),
		DatabasePort:     getEnv("DATABASE_PORT"),
		DatabaseSslMode:  getEnv("DATABASE_SSL_MODE"),
		DatabaseName:     getEnv("DATABASE_NAME"),
		Sslrootcert:      getEnv("SSLROOTCERT"),
		ServerPort:       ":" + getEnv("SERVER_PORT"),
		JWTSecret:        getEnv("JWT_SECRET"),
		AllowedOrigins:   getEnv("ALLOWED_ORIGINS"),
	}
}

func confLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	log.Error().Msg("Environment variable " + key + " is not set.")
	panic("Environment variable " + key + " is not set.")
}
