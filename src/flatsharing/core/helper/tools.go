package helper

import (
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
)

func GenUlid() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func GetEnv(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}
