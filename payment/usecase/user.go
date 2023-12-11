package usecase

import (
	"log"

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

func (u *userPaymentUsecase) FetchAllPlans() {}
func (u *userPaymentUsecase) CreateSubscription() {
	client := razorpay.NewClient("rzp_test_siCMLqIerLB4yZ", "w3W7MyJWfOWjW4LPVDMa2nSr")
	// res, resErr := funcs.CreateAccount(10, "muhammedarif0100@gmail.com", "muhammed arif p", "8137833975")
	// if resErr != nil {
	// 	log.Fatal(resErr.Error())
	// }

	// u.repo.CreateRazorpayAccount(context.Background(), 1, res)

	// fmt.Println(res)

	data := map[string]interface{}{
		"plan_id":         "plan_NA8LHpwdTTCCKf",
		"total_count":     3,
		"customer_id":     "cust_NAE72WkEyXuRPP",
		"quantity":        1,
		"customer_notify": 1,
		"addons": []interface{}{
			map[string]interface{}{
				"item": map[string]interface{}{
					"name":     "Delivery charges",
					"amount":   3000,
					"currency": "INR",
				},
			},
		},
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}

	sub, subErr := client.Subscription.Create(data, nil)
	if subErr != nil {
		log.Fatal(subErr)
	}

	client.Subscription.
		fmt.Println(sub)
}

func (u *userPaymentUsecase) CancelSubscription() {}
func (u *userPaymentUsecase) ChangePlan()         {}
