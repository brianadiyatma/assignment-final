package util

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	salt := 8
	pass := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)

	return string(hash)
}