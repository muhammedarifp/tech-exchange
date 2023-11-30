package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// func init() {
// 	config.InitConfig()
// }

func main() {
	http.HandleFunc("/api/v1/users/", reverseProxy("http://localhost:8000/"))
	http.HandleFunc("/api/v1/users/admins", reverseProxy("http://localhost:8000/"))
	http.ListenAndServe(":8080", nil)
}

func reverseProxy(target string) http.HandlerFunc {
	finalUrl, err := url.Parse(target)
	if err != nil {
		log.Println(err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(finalUrl)

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println("--> ", strings.TrimPrefix(r.URL.Path, "/users"))
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/users")
		//c.Request().Header.Set("X-Real-IP", r)

		proxy.ServeHTTP(w, r)

		// return nil
	}
}
