package main

import (
	"math/rand"

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

//
// UUIDs are too long, use this instead
//
func CreateID(size int) string {
	var possible = []rune("abcdefghijkmnopqrstuvwxyz0123456789-_+*^")

	tmp := make([]rune, size)
	for i := range tmp {
		tmp[i] = possible[rand.Intn(len(possible))]
	}

	return string(tmp)
}
