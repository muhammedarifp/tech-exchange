package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// Auth middle ware here
func AuthUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("Token")
		_, err := jwt.Parse(client, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			w.Write([]byte("Errorrrr !!"))
		} else {
			fmt.Println("Seccuss")
			next.ServeHTTP(w, r)
		}
	})
}
