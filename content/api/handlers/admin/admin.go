package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/usecases/interfaces"
)

type AdminContentHandler struct {
	usecase interfaces.AdminContentUseCase
}

func NewAdminContentHandler(usecase interfaces.AdminContentUseCase) *AdminContentHandler {
	return &AdminContentHandler{
		usecase: usecase,
	}
}

func (h *AdminContentHandler) GetallPosts(c echo.Context) error {
	page := c.QueryParam("page")
	pageInt, strconvErr := strconv.Atoi(page)
	if strconvErr != nil {
		return c.String(400, strconvErr.Error())
	}

	posts, usecasErr := h.usecase.GetallPosts(pageInt)
	if usecasErr != nil {
		return c.String(400, usecasErr.Error()+" : usecase")
	}

	return c.JSON(200, posts)
}
