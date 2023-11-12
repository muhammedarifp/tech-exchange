package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/user/config"
)

// Auth middle ware here
func AuthUserMiddleware(next http.Handler) http.Handler {
	cfg := config.GetConfig()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := r.Header.Get("Token")
		_, err := jwt.Parse(client, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT_SECRET), nil
		})
		if err != nil {
			w.Write([]byte("Errorrrr !!"))
		} else {
			fmt.Println("Seccuss")
			next.ServeHTTP(w, r)
		}
	})
}
