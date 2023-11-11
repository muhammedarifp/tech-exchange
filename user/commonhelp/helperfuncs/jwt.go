package helperfuncs

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(userid uint) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"is_admin": false,
	})

	// fmt.Println(os.Getenv("JWT_SECRET"))
	tokenStr, err := claims.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println(err.Error())
	}

	return tokenStr
}

func GetUserIdFromJwt(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return 0, err
	} else {
		fmt.Println(parsedToken.Claims)
		return 0, nil
	}
}
