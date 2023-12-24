package handlers

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/commonHelp/response"
)

func (u *AdminContentHandler) AddNewTag(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "body read error",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "body read error",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	tagVal, ok := data["tag"].(string)
	if !ok {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Field 'tag' not found or not a string",
			Data:       nil,
			Errors:     "input type error",
		})
	}

	newTagData, repoErr := u.usecase.CreateNewTag(tagVal)
	if repoErr != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Internal server error",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
	}

	return c.JSON(200, response.Response{
		StatusCode: 400,
		Message:    "Success",
		Data:       newTagData,
		Errors:     nil,
	})
}
