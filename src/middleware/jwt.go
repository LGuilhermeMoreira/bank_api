package middleware

import (
	"net/http"
	"strings"

	"github.com/LGuilhermeMoreira/bank_api/src/config"
	"github.com/dgrijalva/jwt-go/v4"
)

type JWT struct {
	secret string
}

// NewJWT returns new JWT.
func NewJWT() *JWT {

	return &JWT{
		secret: config.NewConfig().JwtPassword,
	}
}

func (j *JWT) VerifyAuthToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			msg := "Error on authorization field: missing authorization"
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")

		if len(tokenSplit) < 2 {
			msg := "Error on authorization field"
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		auth := tokenSplit[1]

		claim := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(auth, claim, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		})

		if err != nil {
			msg := "Error to process the jwt token"
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		if !token.Valid {
			msg := "Error invalid token"
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
