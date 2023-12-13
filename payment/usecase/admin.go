package usecase

import (
	"context"
	"log"
	"time"

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

func (u *adminPaymentUsecase) AddPlan() {
	client := razorpay.NewClient("rzp_test_siCMLqIerLB4yZ", "w3W7MyJWfOWjW4LPVDMa2nSr")
	data := map[string]interface{}{
		"period":   "monthly",
		"interval": 1,
		"item": map[string]interface{}{
			"name":        "Monthly subscription plan",
			"amount":      1000,
			"currency":    "INR",
			"description": "Description for the test plan",
		},
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}
	plan, planErr := client.Plan.Create(data, nil)
	if planErr != nil {
		log.Fatalf(planErr.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	u.repo.AddPlan(ctx, plan)
}
func (u *adminPaymentUsecase) RemovePlan() {}
