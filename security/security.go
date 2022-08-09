package security

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

func Hash(password string) ([]byte, error) {
	h := sha512.New()

	h.Write([]byte(password))
	password_hashed := hex.EncodeToString(h.Sum(nil))
	return []byte(password_hashed), nil
}

func VerifyPassword(passwordHash, passwordString string) error {
	a, _ := Hash(passwordString)

	if string(a) != passwordHash {

		return errors.New("invalid password")
	}
	return nil
}
