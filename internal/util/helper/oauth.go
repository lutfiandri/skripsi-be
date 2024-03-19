package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateClientSecret(length int) (string, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	encodedKey := base64.URLEncoding.EncodeToString(key)

	return encodedKey, nil
}
