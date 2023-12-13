package funcs

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/tech-exchange/payments/config"
)

func GetuseridFromJwt(token string) string {
	cfg := config.GetConfig()
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil {
		return ""
	} else {
		cliems := parsedToken.Claims.(jwt.MapClaims)
		userid := cliems["userid"].(string)
		fmt.Println(userid)
		return userid
	}
}
