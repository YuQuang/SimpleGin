package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length*4)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	return string(salt), nil
}

func HashPassword(password string, salt string) string {
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(salt),
		3,
		32*1024,
		4,
		32,
	)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		32*1024,
		3,
		4,
		base64.RawStdEncoding.EncodeToString([]byte(salt)),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded
}

func VerifyPassword(encodedHash string, password string) bool {
	parts := strings.Split(encodedHash, "$")
	if len(parts) < 6 {
		return false
	}
	var salt string = parts[4]

	saltByte, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}
	if HashPassword(password, string(saltByte)) != encodedHash {
		return false
	}

	return true
}
