package flash

import (
	"context"
	"html/template"
	"net/http"
)

func CheckForFlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new data map
		data := make(map[string]interface{})

		// Get the flash message
		if flash := Get(w, r); flash != nil {
			data["FlashType"] = flash.Type
			data["FlashMessage"] = template.HTML(flash.Message)
		}

		// Create a new context with the data map
		ctx := context.WithValue(r.Context(), "data", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
