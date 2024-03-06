package debug

import (
	"log"
	"os"
)

// Logger is a simple wrapper for log.Println - Debugging purposes
func Print(input any) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println(input)
}
