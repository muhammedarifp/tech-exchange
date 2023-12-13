package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Token")
		cfg := config.GetConfig()
		if token == "" {
			ctx.JSON(http.StatusBadRequest, response.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Oops! It looks like you're missing your authorization token. Please provide your token to continue.",
				Data:       nil,
				Errors:     "Empty token",
			})
			return
		}

		parsedToken, parseErr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT_SECRET), nil
		})

		if parseErr != nil {
			ctx.JSON(http.StatusBadRequest, response.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Error parsing authorization token. Please ensure your token is valid and try again.",
				Data:       nil,
				Errors:     parseErr,
			})
			return
		}

		if !parsedToken.Valid {
			ctx.JSON(http.StatusBadRequest, response.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid authorization token. Please provide a valid token and try again.",
				Data:       nil,
				Errors:     parseErr,
			})
			return
		}

		ctx.Next()
	}
}
