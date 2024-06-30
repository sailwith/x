package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt passwords using the bcrypt algorithm.
func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify the password is correct.
func IsValid(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
