package logger

import (
	"fmt"
	"time"
)

type VerboseLogger struct{}

func (l VerboseLogger) Info(message string) {
	fmt.Printf("%s:info: %s\n", time.Now(), message)
}

func (l VerboseLogger) Error(errorMessage string) {
	fmt.Printf("%s:ERROR: %s\n", time.Now(), errorMessage)
}
