package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) string {
	salt := []byte("static_salt")
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	encoded := base64.RawStdEncoding.EncodeToString(hash)
	return encoded
}
