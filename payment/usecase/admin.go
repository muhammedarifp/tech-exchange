package usecase

import (
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	usecase "github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type adminPaymentUsecase struct {
	repo interfaces.AdminPaymentRepo
}

func NewAdminPaymentsUsecase(repo interfaces.AdminPaymentRepo) usecase.AdminPaymentUsecase {
	return &adminPaymentUsecase{
		repo: repo,
	}
}

func (u *adminPaymentUsecase) AddPlan()    {}
func (u *adminPaymentUsecase) RemovePlan() {}
