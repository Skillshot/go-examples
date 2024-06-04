package main

import (
	"go-examples/logger"
	"testing"
)

func TestGettingVerboseLoggerIfVerbosityIsGreaterThanZero(t *testing.T) {
	// Arrange
	verbosity := 1
	// Act
	l := getLoggerBasedOnUserVerbosity(verbosity)
	// Assert
	_, ok := l.(logger.VerboseLogger)
	if !ok {
		t.Errorf("want %T, got %T", logger.VerboseLogger{}, l)
	}
}

func Test_getLoggerBasedOnUserVerbosity(t *testing.T) {
	t.Run("Should return VerboseLogger when verbosity is greater than 0", func(t *testing.T) {
		// Arrange
		verbosity := 1
		// Act
		l := getLoggerBasedOnUserVerbosity(verbosity)
		// Assert
		_, ok := l.(logger.VerboseLogger)
		if !ok {
			t.Errorf("want %T, got %T", logger.VerboseLogger{}, l)
		}
	})

	t.Run("Should return StandardLogger when verbosity is 0", func(t *testing.T) {
		// Arrange
		verbosity := 0
		// Act
		l := getLoggerBasedOnUserVerbosity(verbosity)
		// Assert
		_, ok := l.(logger.StandardLogger)
		if !ok {
			t.Errorf("want %T, got %T", logger.StandardLogger{}, l)
		}
	})
}
