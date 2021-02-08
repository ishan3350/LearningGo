package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "Password123!"
	hash, _ := bcrypt.GenerateFromPassword([]byte(s), 10)
	fmt.Println(string(hash))

	// checking if plaintext is matching with hash or not (Password validation method)
	check := bcrypt.CompareHashAndPassword(hash, []byte(s))

	if check != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Success")
	}
}
