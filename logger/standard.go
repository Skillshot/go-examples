package logger

import (
	"fmt"
)

type StandardLogger struct{}

func (l StandardLogger) Info(message string) {
	fmt.Printf("info: %s\n", message)
}

func (l StandardLogger) Error(errorMessage string) {
	fmt.Printf("ERROR: %s\n", errorMessage)
}
