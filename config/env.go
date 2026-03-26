package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env não encontrado")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
