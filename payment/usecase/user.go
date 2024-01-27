package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/config"
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	usecase "github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
	"github.com/razorpay/razorpay-go"
)

type userPaymentUsecase struct {
	repo interfaces.UserPaymentRepo
}

func NewUserPaymentUsecase(repo interfaces.UserPaymentRepo) usecase.UserPaymentUsecase {
	return &userPaymentUsecase{
		repo: repo,
	}
}

func (u *userPaymentUsecase) FetchAllPlans(page int) ([]response.Plans, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	plans, repoErr := u.repo.FetchAllPlans(ctx, page)

	if repoErr != nil {
		return []response.Plans{}, repoErr
	}

	return plans, nil
}

// Create subscription
func (u *userPaymentUsecase) CreateSubscription(userid uint, request request.CreateSubscriptionReq) (response.Subscription, error) {
	cfg := config.GetConfig()

	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	data := map[string]interface{}{
		"plan_id":         request.PlanID,
		"total_count":     24,
		"quantity":        1,
		"customer_notify": 1,
	}

	sub, subErr := client.Subscription.Create(data, nil)
	if subErr != nil {
		return response.Subscription{}, subErr
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	subscription, repoErr := u.repo.CreateSubscription(ctx, sub)
	if repoErr != nil {
		return subscription, repoErr
	}

	return subscription, nil
}

// Cancel subscription
// TODO: Add cancel subscription in repository
func (u *userPaymentUsecase) CancelSubscription(subid string) (response.Subscription, error) {
	cfg := config.GetConfig()
	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	_, err := client.Subscription.Cancel(subid, nil, nil)
	if err != nil {
		return response.Subscription{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	subsc, repoErr := u.repo.CancelSubscription(ctx, subid)
	if repoErr != nil {
		return response.Subscription{}, repoErr
	}

	return subsc, nil
}

func (u *userPaymentUsecase) ChangePlan() {}

func (u *userPaymentUsecase) VerifyPayment(payload, signature string) (bool, error) {
	secret := config.GetConfig().RAZORPAY_SEC
	secretBytes := []byte(secret)

	hash := hmac.New(sha256.New, secretBytes)

	hash.Write([]byte(payload))

	calculatedSignature := hex.EncodeToString(hash.Sum(nil))

	if calculatedSignature != signature {
		return false, errors.New("Signature mismatched")
	}

	return true, nil
}
