package main

import (
	"github.com/google/uuid"
)

/**
 * Get a new rfc4122 universal unique
 * identifier
 */
func CreateUUID() string {
	id := uuid.New()
	return id.String()
}
