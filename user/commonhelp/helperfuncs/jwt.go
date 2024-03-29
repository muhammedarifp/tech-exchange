package helperfuncs

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/user/config"
)

func CreateJwtToken(userid uint, is_admin bool) string {
	cfg := config.GetConfig()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   strconv.Itoa(int(userid)),
		"is_admin": is_admin,
	})

	tokenStr, err := claims.SignedString([]byte(cfg.JWT_SECRET))

	if err != nil {
		fmt.Println(err.Error())
	}

	return tokenStr
}

func GetUserIdFromJwt(token string) (string, error) {
	cfg := config.GetConfig()
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil {
		return "", err
	} else {
		cliems := parsedToken.Claims.(jwt.MapClaims)
		userid := cliems["userid"].(string)
		return userid, nil
	}
}

func IsAdmin(token string) bool {
	cfg := config.GetConfig()
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil {
		return false
	} else {
		cliems := parsedToken.Claims.(jwt.MapClaims)
		is_admin, ok := cliems["is_admin"].(bool)
		if !ok {
			return false
		}
		return is_admin
	}
}
