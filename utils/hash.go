package utils

import "golang.org/x/crypto/bcrypt"

func HashData(data string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), cost)

	return string(bytes), err
}

func CheckData(data string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data))

	return err == nil
}
