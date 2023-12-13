package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/content/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/funcs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type UserPaymentHandler struct {
	usecase interfaces.UserPaymentUsecase
}

func NewUserPaymentHandler(usecase interfaces.UserPaymentUsecase) *UserPaymentHandler {
	return &UserPaymentHandler{
		usecase: usecase,
	}
}

func (a *UserPaymentHandler) FetchPlans(c *gin.Context) {}

// Create subscription handler
func (a *UserPaymentHandler) CreateSubscription(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	useridStr := funcs.GetuseridFromJwt(token)

	fmt.Println("id == ", useridStr)

	if useridStr == "" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user ID in the token. Please provide a valid user ID.",
			Data:       nil,
			Errors:     "Missing or incorrect user ID in the token",
		})
		return
	}

	userid, convErr := strconv.Atoi(useridStr)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user ID in the token. Please provide a valid user ID.",
			Data:       nil,
			Errors:     "Failed to convert user ID to integer",
		})
		return
	}

	body, bodyErr := io.ReadAll(c.Request.Body)
	if bodyErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body. Please provide a valid JSON body.",
			Data:       nil,
			Errors:     bodyErr.Error(),
		})
		return
	}

	var subscriptionReq request.CreateSubscriptionReq
	if err := json.Unmarshal(body, &subscriptionReq); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body. Please provide a valid JSON body.",
			Data:       nil,
			Errors:     "Failed to parse JSON body: " + err.Error(),
		})
		return
	}

	sub, subErr := a.usecase.CreateSubscription(uint(userid), subscriptionReq)
	if subErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Internal server error.",
			Data:       nil,
			Errors:     subErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       sub,
		Errors:     nil,
	})
}
func (a *UserPaymentHandler) CancelSubscription(c *gin.Context) {}
func (a *UserPaymentHandler) ChangePlan(c *gin.Context)         {}
