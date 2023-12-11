package helperfuncs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/tech-exchage/gateway/config"
	"github.com/pkg/errors"
)

func ValidateUserAuthToken(token string, cfg config.Config) (bool, error) {
	_, parsedErr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if parsedErr != nil {
		return false, errors.Wrap(parsedErr, "token parse error")
	}

	return true, nil
}
