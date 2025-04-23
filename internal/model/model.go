package model

import (
	"AetherGo/internal/log"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Errorf("Failed to encrypt password: %v", err)
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Errorf("Failed to compare password: %v", err)
		return false
	}
	return true
}
