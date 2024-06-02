package utility

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetHesh(pass string) (string, error) {
	fmt.Println("here in gethesh 1")
	hash, er := bcrypt.GenerateFromPassword([]byte(pass), 14)
	fmt.Println("here in gethesh 2")
	return string(hash), er
}

func Compare(pass string, hash string) bool {
	er := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	fmt.Println("Her in compare ", er)
	return er == nil

}
