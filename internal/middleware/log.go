package middleware

import (
	"fmt"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("======================\n")
		fmt.Printf("Method: %v\n", r.Method)
		fmt.Printf("Path: %v\n", r.URL.Path)
		fmt.Printf("======================\n\n")

		// call the next handle
		next.ServeHTTP(w, r)
	})
}

// func GetJWT()
