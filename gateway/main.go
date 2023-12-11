package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/muhammedarifp/tech-exchage/gateway/config"
)

var (
	cfg_pub config.Config
)

func init() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("load configuartion error : %v", err)
	}

	cfg_pub = *cfg
}

func main() {
	// User
	http.HandleFunc(cfg_pub.USER_COMMON, reverseProxy(cfg_pub.USER_SERVICE))
	http.HandleFunc(cfg_pub.USER_ADMIN_COMMON, reverseProxy(cfg_pub.USER_SERVICE))

	// Content
	http.HandleFunc(cfg_pub.CONTENT_COMMON, reverseProxy(cfg_pub.CONTENT_SERVICE))
	http.HandleFunc(cfg_pub.CONTENT_ADMIN_COMMON, reverseProxy(cfg_pub.CONTENT_SERVICE))

	// Notification
	http.HandleFunc(cfg_pub.NOTIFICATION_COMMON, reverseProxy(cfg_pub.NOTIFICATION_SERVICE))

	// Payment
	http.HandleFunc(cfg_pub.PAYMENT_COMMON, reverseProxy(cfg_pub.PAYMENT_SERVICE))
	http.HandleFunc(cfg_pub.PAYMENT_ADMIN_COMMON, reverseProxy(cfg_pub.PAYMENT_SERVICE))

	// Run server
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
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/users")
		//c.Request().Header.Set("X-Real-IP", r)

		proxy.ServeHTTP(w, r)

		// return nil
	}
}

// func authUserEndpoints(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 	})
// }
