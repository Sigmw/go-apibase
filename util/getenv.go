package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(variable string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading dotenv file: %v", err)
	}
	key := os.Getenv(variable)
	if key == "" {
		return key, fmt.Errorf("variable %v does not exist", variable)
	}
	return key, nil
}
