package utils

import (
	"log"
	"os"
)

// Logger is a simple wrapper for log.Println - Debugging purposes
func Logger(input string) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println(input)
}
