package service

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type AuthService struct {
}

func (as *AuthService) Login(
	email string,
	username string,
	password string,
) (string, error) {
	if email == "" {

	}
	if username == "" {

	}

	hash := argon2.IDKey(
		[]byte(password),
		[]byte("3f347f9643deb1679d43765d4d38cae8"),
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
		base64.RawStdEncoding.EncodeToString([]byte("3f347f9643deb1679d43765d4d38cae8")),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return string(encoded), nil
}
