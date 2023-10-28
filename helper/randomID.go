package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateID() string {
	bytes := make([]byte, 5)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)[:5]
}
