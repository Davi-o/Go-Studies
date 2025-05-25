package sorting

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingWithBCrypt() {
	password := "anyPassword"
	wrongPassword := "anotherPassword"
	sliceOfBytes, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		fmt.Println(error)
	}

	checkPassHash := bcrypt.CompareHashAndPassword(sliceOfBytes, []byte(password))
	if checkPassHash == nil {
		fmt.Println(string(sliceOfBytes))
	}

	checkPassHash = bcrypt.CompareHashAndPassword(sliceOfBytes, []byte(wrongPassword))
	if checkPassHash != nil {
		fmt.Println(checkPassHash)
	}

}

