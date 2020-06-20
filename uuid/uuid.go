package uuid

import (
	"github.com/google/uuid"
)

// Get returns a new randomly-generated UUID value.
func Get() string {
	uuid := uuid.New()

	return uuid.String()
}
