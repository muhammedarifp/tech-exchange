package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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

func (u *ContentUserHandler) LikePost(c echo.Context) error {
	postid := c.QueryParam("post_id")
	if postid == "" {
		return c.String(400, "Your postid is empty")
	}

	_, err := u.usecase.LikePost(postid)
	if err != nil {
		return c.String(400, err.Error())
	}

	return c.String(200, "Liked")
}
