package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/funcs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
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

func (a *UserPaymentHandler) FetchPlans(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     "invalid page number",
		})
		return
	}
	pageInt, convErr := strconv.Atoi(page)
	if convErr != nil {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     "invalid page number",
		})
		return
	}
	plans, err := a.usecase.FetchAllPlans(pageInt)
	if err != nil {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     err,
		})
		return
	}

	c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       plans,
		Errors:     nil,
	})
}

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
func (a *UserPaymentHandler) CancelSubscription(c *gin.Context) {
	subid := c.Query("subscription")
	if subid == "" {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "invalid subscription id",
			Data:       nil,
			Errors:     "invalid id provided",
		})
		return
	}

	resp, repoErr := a.usecase.CancelSubscription(subid)
	if repoErr != nil {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
		return
	}

	c.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       resp,
		Errors:     nil,
	})
}
func (a *UserPaymentHandler) ChangePlan(c *gin.Context) {}

func (a *UserPaymentHandler) VerifyPayment(c *gin.Context) {
	// token := c.Request.Header.Get("Token")
	// useridStr := funcs.GetuseridFromJwt(token)

	var userData map[string]interface{}

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

	json.Unmarshal(body, &userData)

	payload, ok1 := userData["payload"].(string)
	signature, ok2 := userData["signature"].(string)

	if !ok1 || !ok2 {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body. Please provide a valid JSON body.",
			Data:       nil,
			Errors:     "Payload or signature is missing",
		})
		return
	}

	if payload == "" || signature == "" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body. Please provide a valid JSON body.",
			Data:       nil,
			Errors:     "Payload or signature is missing",
		})
		return
	}

	_, repoErr := a.usecase.VerifyPayment(payload, signature)
	if repoErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Internal server error.",
			Data:       nil,
			Errors:     repoErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       "Payment verified successfully",
		Errors:     nil,
	})

}
