package helper

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Logger is the global variable container logger
var Logger *zap.SugaredLogger

// InitLogger init Logger
func InitLogger() {
	log, err := zap.NewProduction()
	if err != nil {
		fmt.Print("Can't start logger")
		os.Exit(-1)
	}
	Logger = log.Sugar()
}
