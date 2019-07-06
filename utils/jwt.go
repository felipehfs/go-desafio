package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	SECRET_KEY = []byte("s3cret")
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// EnabledJwt is a middleware that checks the existent token
func EnabledJwt() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (interface{}, error) {
					return SECRET_KEY, nil
				})

			if err == nil {
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprint(w, "Token não é válido")
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				errorMessage := map[string]string{
					"status":  "error",
					"message": "Usuário não pode ser autenticado!",
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(errorMessage)
			}
		}
	}
}
