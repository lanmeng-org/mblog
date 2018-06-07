package util

import (
	"github.com/elithrar/simple-scrypt"
)

func EncryptPwd(password []byte) ([]byte, error) {
	return scrypt.GenerateFromPassword(password, scrypt.DefaultParams)
}

func CheckPwd(password, hash []byte) bool {
	err := scrypt.CompareHashAndPassword(hash, password)

	return err == nil
}