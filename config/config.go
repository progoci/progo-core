package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config is the configuration handler.
type Config struct{}

// New loads a .env file.
func New(filepath string) (*Config, error) {
	err := godotenv.Overload(filepath)

	if err != nil {
		return nil, err
	}

	return &Config{}, nil
}

// Get retrieves a environment variable.
func (Config) Get(key string) string {
	return os.Getenv(key)
}
