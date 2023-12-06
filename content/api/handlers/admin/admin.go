package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/commonHelp/jwt"
	"github.com/muhammedarifp/content/commonHelp/response"
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

func (u *AdminContentHandler) DeleteContent(c echo.Context) error {
	// Extract user ID from the JWT token
	userID := jwt.GetuseridFromJwt(c.Request().Header.Get("Token"))

	// Check if the user ID is invalid
	if userID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user ID",
			Data:       nil,
			Errors:     []string{"Invalid user ID"},
		})
	}

	// Extract post ID from the URL parameter
	postID := c.QueryParam("postid")

	// Check if the post ID is invalid
	if postID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid post ID",
			Data:       nil,
			Errors:     []string{"Invalid post ID"},
		})
	}

	// Remove the post
	content, repoErr := u.usecase.RemovePost(postID, userID)

	// Check for errors during post removal
	if repoErr != nil {
		return c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to remove post",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Post removed successfully",
		Data:       content,
		Errors:     nil,
	})
}
