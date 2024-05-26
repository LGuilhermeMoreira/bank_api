package middleware

import (
	"fmt"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		// call the next handle
		next.ServeHTTP(w, r)
	})
}

// func GetJWT()
