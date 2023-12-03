package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/commonHelp/jwt"
	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/commonHelp/response"
)

func AuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Token")
		if !jwt.VerifyUser(token) {
			return c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "0unautheraized user",
				Data:       nil,
				Errors:     "permission denaid",
			})
		}

		req, _ := http.NewRequest("GET", "http://192.168.193.17:8080/api/v1/users/account", nil)
		req.Header.Set("Token", token)
		client := http.Client{}
		user, user_err := client.Do(req)
		if user_err != nil {
			return c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "1unautheraized user",
				Data:       nil,
				Errors:     user_err.Error(),
			})
		}
		userValbyte, readErr := io.ReadAll(user.Body)
		if readErr != nil {
			log.Println(readErr.Error())
			return nil
		}
		fmt.Println(string(userValbyte))
		var userVal requests.UserValue
		if err := json.Unmarshal(userValbyte, &userVal); err != nil {
			return c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "unmarshel error found",
				Data:       nil,
				Errors:     err.Error(),
			})
		}

		//
		if userVal.Data.Is_banned {
			return c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "10unautheraized user",
				Data:       nil,
				Errors:     "account banned",
			})
		}

		if !userVal.Data.Is_active {
			return c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "11unautheraized user",
				Data:       nil,
				Errors:     "account deactivated",
			})
		}

		return next(c)
	}
}
