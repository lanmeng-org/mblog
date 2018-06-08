package util

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func CheckPwd(hash, password[]byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}