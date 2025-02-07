package utils

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvLoader struct{}

func (e *EnvLoader) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
