package usecase

import (
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	usecase "github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type userPaymentUsecase struct {
	repo interfaces.UserPaymentRepo
}

func NewUserPaymentUsecase(repo interfaces.UserPaymentRepo) usecase.UserPaymentUsecase {
	return &userPaymentUsecase{
		repo: repo,
	}
}

func (u *userPaymentUsecase) FetchAllPlans()      {}
func (u *userPaymentUsecase) CreateSubscription() {}
func (u *userPaymentUsecase) CancelSubscription() {}
func (u *userPaymentUsecase) ChangePlan()         {}
