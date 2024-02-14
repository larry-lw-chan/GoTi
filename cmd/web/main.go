package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/router"
)

func main() {
	// Get Routes
	r := router.Routes()

	// Start Server
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
