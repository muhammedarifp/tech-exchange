package usecase

import (
	"context"
	"log"
	"time"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/config"
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	usecase "github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
	"github.com/razorpay/razorpay-go"
)

type adminPaymentUsecase struct {
	repo interfaces.AdminPaymentRepo
}

func NewAdminPaymentsUsecase(repo interfaces.AdminPaymentRepo) usecase.AdminPaymentUsecase {
	return &adminPaymentUsecase{
		repo: repo,
	}
}

func (u *adminPaymentUsecase) AddPlan(enterData request.Plans) (response.Plans, error) {
	cfg := config.GetConfig()
	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	data := map[string]interface{}{
		"period":   enterData.Period,
		"interval": enterData.Interval,
		"item": map[string]interface{}{
			"name":        enterData.Name,
			"amount":      enterData.Amount * 100,
			"currency":    "INR",
			"description": enterData.Description,
		},
	}
	plan, planErr := client.Plan.Create(data, nil)
	if planErr != nil {
		log.Fatalf(planErr.Error())
		return response.Plans{}, planErr
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := u.repo.AddPlan(ctx, plan)
	if err != nil {
		return response.Plans{}, err
	}

	return resp, nil
}

func (u *adminPaymentUsecase) RemovePlan(planid string) (response.Plans, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	plan, err := u.repo.RemovePlan(ctx, planid)
	return plan, err
}
