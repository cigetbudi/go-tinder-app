package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal membaca file .env ", err)
	}
	return os.Getenv(key)
}
