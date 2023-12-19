package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/funcs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/msgs"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/config"
	"github.com/muhammedarifp/tech-exchange/payments/rabbitmq"
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
	// cfg := config.GetConfig()
	// client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	// offset := (page - 1) * 10
	// options := map[string]interface{}{
	// 	"count": 10,
	// 	"skip":  offset,
	// }
	// plans, planErr := client.Plan.All(options, nil)
	// if planErr != nil {
	// 	fmt.Println("Error found :  ", planErr)
	// 	return map[string]interface{}{}, planErr
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	plans, repoErr := u.repo.FetchAllPlans(ctx, page)

	if repoErr != nil {
		return []response.Plans{}, repoErr
	}

	return plans, nil
}

func (u *userPaymentUsecase) CreateSubscription(userid uint, request request.CreateSubscriptionReq) (response.Subscription, error) {
	cfg := config.GetConfig()

	account, accountErr := u.repo.FetchRazorpayAccount(userid)
	if accountErr != nil {
		return response.Subscription{}, accountErr
	}
	if account.UserID == 0 {
		return response.Subscription{}, errors.New("User account not found")
	}

	fmt.Println("account : ", account)

	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	data := map[string]interface{}{
		"plan_id":         request.PlanID,
		"total_count":     24,
		"customer_id":     account.RazorpayID,
		"quantity":        1,
		"customer_notify": 1,
		"addons": []interface{}{
			map[string]interface{}{
				"item": map[string]interface{}{
					"name":     "Premium",
					"amount":   6000,
					"currency": "INR",
				},
			},
		},
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

	return subscription, repoErr
	// return response.Subscription{}, nil
}

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

func StartMessageServer() {
	u := userPaymentUsecase{}
	u.CreatePaymentAccount()
}

func (u *userPaymentUsecase) CreatePaymentAccount() {
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
				fmt.Println(".............................", err.Error())
			}

			fmt.Println(string(msg.Body))

			accountVal, err := funcs.CreateAccount(req.Email, req.Name)
			if err != nil {
				fmt.Println(err)
				return
			}

			u.repo.CreateRazorpayAccount(context.Background(), req.UserID, accountVal)
		}
	}()

	select {}
}

// func (r *userPaymentUsecase) Alllll() {
// 	optional := map[string]interface{}{
// 		"count": 10,
// 	}
// 	cfg := config.GetConfig()
// 	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
// 	body, err := client.Customer.All(optional, nil)
// 	if err != nil {
// 		return
// 	}
// 	var c msgs.Collection
// 	marshelData, errr := json.Marshal(body)
// 	if errr != nil {
// 		log.Fatal(errr)
// 	}

// 	if err := json.Unmarshal(marshelData, &c); err != nil {
// 		log.Fatal(err)
// 	}

// 	repository.Temparory(c)
// }
