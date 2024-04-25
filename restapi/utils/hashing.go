package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Error comparing password: %v", hashedPassword)
		log.Printf("Error comparing password: %v", err)
	}
	return err == nil
}
