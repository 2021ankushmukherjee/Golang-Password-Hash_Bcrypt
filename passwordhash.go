package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	password := "ankushmukherjee"
	hash, _ := HashPassowrd(password)

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Matching status ", match)
}
