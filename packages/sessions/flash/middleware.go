package flash

import (
	"context"
	"log"
	"net/http"
)

func TryGetFlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new data map
		data := make(map[string]interface{})

		// Get the flash message
		flash := Get(w, r)
		log.Println(flash)
		if flash != nil {
			data["Flash"] = flash
		}

		// Create a new context with the data map
		ctx := context.WithValue(r.Context(), "data", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
