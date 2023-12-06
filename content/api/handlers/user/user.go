package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/muhammedarifp/content/commonHelp/jwt"
	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/commonHelp/response"
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

// @Summary Create new post
// @Description Signup new user
// @Tags Content / User
// @Accept json
// @Produce json
// @Param user body requests.CreateNewPostRequest true "User information for signup"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/create-post [post]
// @BasePath /api/v1/users
func (u *ContentUserHandler) CreateNewPost(c echo.Context) error {
	userid := jwt.GetuseridFromJwt(c.Request().Header.Get("Token"))
	if userid == "" {
		fmt.Println("Userid is nill")
		return nil
	}

	userBody, bodyErr := io.ReadAll(c.Request().Body)
	if bodyErr != nil {
		fmt.Println("Body errrr")
		return nil
	}

	var userPost requests.CreateNewPostRequest
	if err := json.Unmarshal(userBody, &userPost); err != nil {
		return c.String(400, err.Error())
	}

	val, err := u.usecase.CreatePost(userid, userPost)
	if err != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "failure",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	return c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       val,
		Errors:     nil,
	})
}

// @Summary Create new comment
// @Description add owm comment on any post
// @Tags Content / User
// @Accept json
// @Produce json
// @Param user body requests.CreateNewPostRequest true "User information for signup"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/comment [post]
// @BasePath /api/v1/users
func (u *ContentUserHandler) CreateComment(c echo.Context) error {
	postid := c.QueryParam("post_id")
	if postid == "" {
		return c.String(400, "invalid postid")
	}
	userid := jwt.GetuseridFromJwt(c.Request().Header.Get("Token"))
	if userid == "" {
		return c.String(400, "invalid userid")
	}

	var requestBody map[string]interface{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		return c.String(400, err.Error())
	}

	commentText := requestBody["comment"].(string)
	text := strings.TrimSpace(commentText)
	if text == "" {
		return c.String(400, "Your comment text is empty, empty comment not allowed !")
	}

	u.usecase.CreateComment(userid, postid, commentText)

	return c.String(200, fmt.Sprintf("postid is %s", postid))
}

// @Summary Like post
// @Description Like post
// @Tags Content / User
// @Accept json
// @Produce json
// @Param param query string true "Example parameter"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/like [post]
// @BasePath /api/v1/users
func (u *ContentUserHandler) LikePost(c echo.Context) error {
	postid := c.QueryParam("post_id")
	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Invalid token provided",
			Data:       nil,
			Errors:     "invalid auth token",
		})
	}
	userid := jwt.GetuseridFromJwt(token)
	if postid == "" {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Invalid params, Post ID is required",
			Data:       nil,
			Errors:     "invalid params",
		})
	}

	content, err := u.usecase.LikePost(postid, userid)
	if err != nil {
		return c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Internal server error, Try again later",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	return c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "Hooray! You've successfully liked the post!",
		Data:       content,
		Errors:     nil,
	})
}

// @Summary Update content
// @Description Like post
// @Tags Content / User
// @Accept json
// @Produce json
// @Param param query string true "Example parameter"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/update [put]
// @BasePath /api/v1/users
func (u *ContentUserHandler) UpdateContent(c echo.Context) error {
	// Extract token from the Authorization header
	token := c.Request().Header.Get(echo.HeaderAuthorization)
	userID := jwt.GetuseridFromJwt(token)

	if userID == "" {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "User ID fetch error",
			Data:       nil,
			Errors:     []string{"User ID fetch error"},
		})
	}

	// Read the request body
	ioBodyVal, ioErr := io.ReadAll(c.Request().Body)
	if ioErr != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "IO error reading request body",
			Data:       nil,
			Errors:     []string{ioErr.Error()},
		})
	}

	// Unmarshal JSON
	var userPost requests.UpdatePostRequest
	if err := json.Unmarshal(ioBodyVal, &userPost); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "JSON unmarshaling error",
			Data:       nil,
			Errors:     []string{err.Error()},
		})
	}

	// Update post
	content, usecaseErr := u.usecase.UpdatePost(userPost, userID)
	if usecaseErr != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Update post error",
			Data:       nil,
			Errors:     []string{usecaseErr.Error()},
		})
	}

	// Return updated content
	return c.JSON(http.StatusOK, content)
}

// @Summary Delete post
// @Description Like post
// @Tags Content / User
// @Accept json
// @Produce json
// @Param param query string true "Example parameter"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/delete [delete]
// @BasePath /api/v1/users
func (u *ContentUserHandler) DeleteContent(c echo.Context) error {
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

// @Summary Delete post
// @Description Like post
// @Tags Content / User
// @Accept json
// @Produce json
// @Param param query string true "Example parameter"
// @Success 200 {object} response.Response "User created success"
// @Failure 400 {object} response.Response "Bad Request"
// @Router /api/v1/contents/delete [delete]
// @BasePath /api/v1/users
func (u *ContentUserHandler) GetUserContents(c echo.Context) error {
	token := c.Request().Header.Get("Token")
	userid := jwt.GetuseridFromJwt(token)
	if userid == "" {
		//
	}

	page := c.QueryParam("page")
	pageInt, strconverr := strconv.Atoi(page)
	if strconverr != nil {
		return c.String(400, strconverr.Error())
	}

	contents, repoErr := u.usecase.GetUserPosts(userid, pageInt)
	if repoErr != nil {
		return c.String(400, "repo eeeee")
	}

	return c.JSON(200, contents)
}

func (h *ContentUserHandler) GetallPosts(c echo.Context) error {
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
