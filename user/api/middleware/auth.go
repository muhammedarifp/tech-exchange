package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/repository"
)

// Auth middle ware here
func AuthUserMiddleware(next http.Handler) http.Handler {
	cfg := config.GetConfig()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		client := r.Header.Get("Token")
		_, err := jwt.Parse(client, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT_SECRET), nil
		})
		if err != nil {
			marshelResp, marshelErr := json.Marshal(response.Response{
				StatusCode: 400,
				Message:    "Permission Denied",
				Data:       nil,
				Errors:     "Invalid Authentication Token Provided",
			})
			if marshelErr != nil {
				log.Println(marshelErr.Error())
			}
			w.WriteHeader(400)
			w.Write(marshelResp)
			return
		} else {
			userid, tokenErr := helperfuncs.GetUserIdFromJwt(client)
			if tokenErr != nil {
				marshelResp, marshelErr := json.Marshal(response.Response{
					StatusCode: 400,
					Message:    "Permission Denied",
					Data:       nil,
					Errors:     tokenErr.Error(),
				})
				if marshelErr != nil {
					log.Println(marshelErr.Error())
				}
				w.WriteHeader(400)
				w.Write(marshelResp)
				return
			}

			//
			userVal, repoErr := repository.FetchUserUsingID_public(userid)

			if repoErr != nil {
				marshelResp, marshelErr := json.Marshal(response.Response{
					StatusCode: 400,
					Message:    "Permission Denied",
					Data:       nil,
					Errors:     repoErr.Error(),
				})
				if marshelErr != nil {
					log.Println(marshelErr.Error())
				}
				w.WriteHeader(400)
				w.Write(marshelResp)
				return
			}

			//
			if userVal.Is_banned {
				marshelResp, marshelErr := json.Marshal(response.Response{
					StatusCode: 400,
					Message:    "Permission Denied",
					Data:       nil,
					Errors:     "User account is banned",
				})
				if marshelErr != nil {
					log.Println(marshelErr.Error())
				}
				w.WriteHeader(400)
				w.Write(marshelResp)
				return
			}

			if !userVal.Is_active {
				// user account deleted
				marshelResp, marshelErr := json.Marshal(response.Response{
					StatusCode: 400,
					Message:    "Permission Denied",
					Data:       nil,
					Errors:     "User account is deactivated",
				})
				if marshelErr != nil {
					log.Println(marshelErr.Error())
				}
				w.WriteHeader(400)
				w.Write(marshelResp)
				return
			}

			next.ServeHTTP(w, r)
		}
	})
}
