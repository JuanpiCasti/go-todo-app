package config

import (
	"log"
	"os"
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
}

var CFG *Config

func LoadConfig() {
	CFG = &Config{
		DatabaseUser:     getEnv("DATABASE_USER"),
		DatabasePassword: getEnv("DATABASE_PASSWORD"),
		DatabaseHost:     getEnv("DATABASE_HOST"),
		DatabasePort:     getEnv("DATABASE_PORT"),
		DatabaseSslMode:  getEnv("DATABASE_SSL_MODE"),
		DatabaseName:     getEnv("DATABASE_NAME"),
		Sslrootcert:      getEnv("SSLROOTCERT"),
		ServerPort:       getEnv("SERVER_PORT"),
	}
}

func getEnv(key string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	log.Fatal("Environment variable " + key + " is not set.")
	panic("Environment variable " + key + " is not set.")
}
