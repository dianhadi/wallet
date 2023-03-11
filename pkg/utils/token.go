package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken() string {
	// Generate a random byte slice of size 21
	randomBytes := make([]byte, 21)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Encode the random bytes as a hexadecimal string
	token := hex.EncodeToString(randomBytes)

	return token
}
