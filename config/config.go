package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env.sample")
	if err != nil {
		fmt.Print("Error loading .env file")
		err = godotenv.Load("../.env.sample")
		if err != nil {
			err = godotenv.Load("../../.env.sample")
			if err != nil {
				err = godotenv.Load("./.env.sample")
			}
		}
	}

	return os.Getenv(key)
}
