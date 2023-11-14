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
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		cfg := config.GetConfig()
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
