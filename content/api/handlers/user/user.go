package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/usecases/interfaces"
)

type ContentUserHandler struct {
	usecase interfaces.ContentUserUsecase
}

func NewContentUserHandler(usecase interfaces.ContentUserUsecase) *ContentUserHandler {
	return &ContentUserHandler{
		usecase: usecase,
	}
}

func (u *ContentUserHandler) CreateNewPost(c echo.Context) error {
	val, err := u.usecase.CreateNewPost()
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"error": nil,
		"data":  val,
	})
}
