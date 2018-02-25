package helper

import "golang.org/x/crypto/bcrypt"

func BcryptGen(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err)
	}
	return string(hash[:])
}
