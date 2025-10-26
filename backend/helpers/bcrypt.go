package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (res string) {
	// tingkat acak/kesulitan pwd
	cost := 10

	// ubah password menjadi byte sliceW
	password := []byte(s)

	// generate password yang sudah menjadi byte, cost / tingkat kesulitan
	// akan menghasilkan hash dalam bentuk byte slice
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		msg := fmt.Sprintf("Error generate pwd %s", err)
		log.Println(msg)
		return
	}

	// ubah byte slice menjadi string dan di assign ke res
	res = string(hash)
	return
}

func ComparePassword(hash, pwd []byte) bool {
	// hash => hash password
	// pwd => plain text password
	err := bcrypt.CompareHashAndPassword(hash, pwd)

	// jika ada err brarti tidak sama
	if err != nil {
		log.Println("Invalid password")
		return false
	}

	// true jika pwd dan hash isinya sama
	return true
}
