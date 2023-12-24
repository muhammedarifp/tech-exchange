package handlers

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/commonHelp/jwt"
	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/commonHelp/response"
)

func (u *ContentUserHandler) Getalltags(c echo.Context) error {
	resp, repoErr := u.usecase.FetchAllTags()
	if repoErr != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "unmarshel error",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
	}

	return c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       resp,
		Errors:     nil,
	})
}

func (u *ContentUserHandler) FollowTag(c echo.Context) error {
	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "body error",
			Data:       nil,
			Errors:     "auth token error",
		})
	}

	userid := jwt.GetuseridFromJwt(token)
	if userid == "" {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "body error",
			Data:       nil,
			Errors:     "invalid userid provided",
		})
	}

	body, bodyErr := io.ReadAll(c.Request().Body)
	if bodyErr != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "body error",
			Data:       nil,
			Errors:     bodyErr.Error(),
		})
	}

	var followTagReq requests.FollowTagReq
	if err := json.Unmarshal(body, &followTagReq); err != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "unmarshel error",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	resp, repoErr := u.usecase.FollowTag(userid, followTagReq)
	if repoErr != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "unmarshel error",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
	}

	return c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       resp,
		Errors:     nil,
	})
}
