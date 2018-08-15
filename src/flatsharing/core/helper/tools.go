package helper

import (
	"bytes"
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
	"golang.org/x/crypto/bcrypt"
)

// BcryptGen take string and return hashed string
func BcryptGen(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err)
	}
	return string(hash[:])
}

// GenUlid generate an ulid
func GenUlid() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// GenToken generate a random token
func GenToken() (string, error) {
	allowCharac := "abcdefghjiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var buffer bytes.Buffer

	for index := 0; index < 128; index++ {
		_, err := buffer.WriteString(string(allowCharac[int(rand.Float32()*float32(len(allowCharac)))]))
		if err != nil {
			return "", err
		}
	}
	return buffer.String(), nil
}

// GetEnv get an environement variable or set with default value if not found
func GetEnv(name, def string) string {
	str := os.Getenv(name)
	if str == "" {
		return def
	}
	return str
}
