package helper

import (
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenUlid generate an ulid
func GenUlid() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// GetEnv get an environement variable or set with default value if not found
func GetEnv(name, def string) string {
	str := os.Getenv(name)
	if str == "" {
		return def
	}
	return str
}

// GenToken generate a random token
func GenToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789" + "!@#$%^&*()"

	b := make([]byte, 128)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
