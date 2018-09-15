package main

import (
	"math/rand"

	"github.com/google/uuid"
)

const characterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString creates a random token string of length n
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterBytes[rand.Intn(len(character_bytes))]
	}
	return string(b)
}



func GetHospitalForID(id uuid.UUID) (Hospital, error) {
	return Hospital{}, nil
}
