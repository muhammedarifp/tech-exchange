package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(nxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("15:04:05") + " | " + r.RequestURI + " | " + r.Method)
		// Call Nxt
		nxt.ServeHTTP(w, r)
	})
}
