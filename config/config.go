package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Init loads a .env file in root directory.
func Init(filepath string) {
	err := godotenv.Overload(filepath)

	if err != nil {
		panic("Could not load .env file" + err.Error())
	}
}

// Get retrieves a environment variable.
func Get(key string) string {

	return os.Getenv(key)
}
