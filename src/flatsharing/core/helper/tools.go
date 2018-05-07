package helper

import (
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
)

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
