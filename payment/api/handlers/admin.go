package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type AdminPaymentHandler struct {
	usecase interfaces.AdminPaymentUsecase
}

func NewAdminPaymentHandler(usecase interfaces.AdminPaymentUsecase) *AdminPaymentHandler {
	return &AdminPaymentHandler{
		usecase: usecase,
	}
}

func (a *AdminPaymentHandler) AddPlan(c *gin.Context) {
	enterData := request.Plans{}
	dataByte, readErr := io.ReadAll(c.Request.Body)
	if readErr != nil {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "boby read error",
			Data:       nil,
			Errors:     readErr.Error(),
		})
		return
	}

	if err := json.Unmarshal(dataByte, &enterData); err != nil {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "unmarshel error",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if enterData.Period == "daily" || enterData.Period == "weekly" || enterData.Period == "monthly" || enterData.Period == "yearly" {

		plan, err := a.usecase.AddPlan(enterData)
		if err != nil {
			c.JSON(400, response.Response{
				StatusCode: 400,
				Message:    "internal server error",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

		c.JSON(200, response.Response{
			StatusCode: 200,
			Message:    "success",
			Data:       plan,
			Errors:     nil,
		})

	} else {
		c.JSON(400, response.Response{
			StatusCode: 400,
			Message:    fmt.Sprintf("your entered perioed is (%s) incorrect. correct that", enterData.Period),
			Data:       nil,
			Errors:     "incorrect input",
		})
		return
	}

	// c.JSON(200, response.Response{
	// 	StatusCode: 200,
	// 	Message:    "success",
	// 	Data:       enterData,
	// 	Errors:     nil,
	// })
}

func (a *AdminPaymentHandler) RemovePlan(c *gin.Context) {
	planid := c.Query("plan")
	if planid == "" {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "input value is incorrect",
			Data:       nil,
			Errors:     "invalid input value",
		})
		return
	}

	plan, repoErr := a.usecase.RemovePlan(planid)
	if repoErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     repoErr.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, response.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       plan,
		Errors:     nil,
	})
}
