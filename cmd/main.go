package main

import (
	"fmt"
	"go-examples/logger"
	"go-examples/server"
)

func main() {
	fmt.Println("Hello, World!")
	consoleLogger := getLoggerBasedOnUserVerbosity(1)
	s := server.New(consoleLogger)
	s.RenderAdminPage("admin")
}

func getLoggerBasedOnUserVerbosity(verbosity int) server.Logger {
	if verbosity > 0 {
		return logger.VerboseLogger{}
	}
	return logger.StandardLogger{}
}
