package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/content/config"
)

func VerifyUser(token string) bool {
	cfg := config.GetConfig()
	if _, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	}); err != nil {
		return false
	} else {
		return true
	}
}
