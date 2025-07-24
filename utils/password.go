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

func VerifyPassword(hashedPassword, inputPassword string) bool {
	salt := []byte("static_salt")
	inputHash := argon2.IDKey([]byte(inputPassword), salt, 1, 64*1024, 4, 32)
	encoded := base64.RawStdEncoding.EncodeToString(inputHash)
	return encoded == hashedPassword
}
