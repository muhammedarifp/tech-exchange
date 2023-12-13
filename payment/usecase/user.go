package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/funcs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/msgs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/rabbitmq"
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

func (u *userPaymentUsecase) FetchAllPlans() {}

func (u *userPaymentUsecase) CreateSubscription(userid uint, request request.CreateSubscriptionReq) (response.Subscription, error) {
	// cfg := config.GetConfig()

	// account, accountErr := u.repo.FetchRazorpayAccount(userid)
	// if accountErr != nil {
	// 	return response.Subscription{}, accountErr
	// }

	// if account.RazorpayID == "" {
	// 	funcs.CreateAccount(userid,)
	// }

	// client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	// data := map[string]interface{}{
	// 	"plan_id":         request.PlanID,
	// 	"total_count":     24,
	// 	"customer_id":     account.RazorpayID,
	// 	"quantity":        1,
	// 	"customer_notify": 1,
	// 	"addons": []interface{}{
	// 		map[string]interface{}{
	// 			"item": map[string]interface{}{
	// 				"name":     "Premium",
	// 				"amount":   6000,
	// 				"currency": "INR",
	// 			},
	// 		},
	// 	},
	// }

	// sub, subErr := client.Subscription.Create(data, nil)
	// if subErr != nil {
	// 	return response.Subscription{}, subErr
	// }

	// fmt.Println("sssssssssssssssssssssuuuuuuuuuuuuuuuuuuuuuuu", sub)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*40)
	// defer cancel()
	// subscription, repoErr := u.repo.CreateSubscription(ctx, sub)
	// if repoErr != nil {
	// 	return subscription, repoErr
	// }

	// return subscription, repoErr

	return response.Subscription{}, nil
}

func (u *userPaymentUsecase) CancelSubscription() {}
func (u *userPaymentUsecase) ChangePlan()         {}

func StartMessageServer() {
	u := userPaymentUsecase{}
	u.CreatePaymentAccount(0)
}

func (u *userPaymentUsecase) CreatePaymentAccount(userid uint) {
	//cfg := config.GetConfig()
	mqconn, mqErr := rabbitmq.CreateRabbitMqConnection()
	if mqErr != nil {
		log.Fatal("Failed to establish connection")
	}

	ch, chErr := mqconn.Channel()
	if chErr != nil {
		log.Fatal("Failed to open channel")
	}

	queue, queErr := ch.QueueDeclare("payment_acc", true, false, false, false, nil)
	if queErr != nil {
		log.Fatal("Failed to declare queue")
	}

	messages, consumeErr := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if consumeErr != nil {
		log.Fatal("Failed to consume messages")
	}

	go func() {
		for msg := range messages {
			req := msgs.PaymentAccount{}
			if err := json.Unmarshal(msg.Body, &req); err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(msg.Body))

			accountVal, err := funcs.CreateAccount(req.UserID, req.Email, req.Name)
			if err != nil {
				fmt.Println(err)
			}
			useridInt, _ := strconv.Atoi(req.UserID)
			u.repo.CreateRazorpayAccount(context.Background(), uint(useridInt), accountVal)
		}
	}()
}
