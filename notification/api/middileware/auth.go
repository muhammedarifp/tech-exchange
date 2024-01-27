package middileware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/notification/config"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if !VerifyUser(token) {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "unautheraized user",
				Data:       nil,
				Errors:     "invalid token",
			})
			c.Abort()
			return
		}

		req, userErr := http.NewRequest("GET", "http://muarif.online/api/v1/users/account", nil)
		if userErr != nil {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "1unautheraized user",
				Data:       nil,
				Errors:     userErr.Error(),
			})
			c.Abort()
			return
		}
		req.Header.Set("Token", token)
		client := http.Client{}
		user, user_err := client.Do(req)
		if user_err != nil {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "",
				Data:       nil,
				Errors:     user_err.Error(),
			})
			c.Abort()
			return
		}
		userValbyte, readErr := io.ReadAll(user.Body)
		if readErr != nil {
			log.Println(readErr.Error())
			return
		}
		fmt.Println(string(userValbyte))
		var userVal requests.UserValue
		if err := json.Unmarshal(userValbyte, &userVal); err != nil {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "",
				Data:       nil,
				Errors:     err.Error(),
			})
			c.Abort()
			return
		}

		//
		if userVal.Data.Is_banned {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "10unautheraized user",
				Data:       nil,
				Errors:     "account banned",
			})
			c.Abort()
			return
		}

		if !userVal.Data.Is_active {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "11unautheraized user",
				Data:       nil,
				Errors:     "account deactivated",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

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
