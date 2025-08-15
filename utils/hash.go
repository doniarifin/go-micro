package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	byt, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	return string(byt), err
}

func CheckHashPassword(hashed, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err == nil
}
