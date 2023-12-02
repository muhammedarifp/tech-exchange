package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	commonhelp "github.com/muhammedarifp/tech-exchange/notification/commonHelp"
	"github.com/muhammedarifp/tech-exchange/notification/rabbitmq"
	interfaces "github.com/muhammedarifp/tech-exchange/notification/repository/interface"
	usercaseInterfaces "github.com/muhammedarifp/tech-exchange/notification/usecase/interfaces"
)

type notificationUsecase struct {
	repo interfaces.NotificationRepo
}

func NewNotificationUseCase(repo interfaces.NotificationRepo) usercaseInterfaces.NotificationUsecase {
	return &notificationUsecase{
		repo: repo,
	}
}

var (
	mu sync.Mutex
)

func (u *notificationUsecase) StoreNotificationsOnDB() {
	conn, connErr := rabbitmq.NewRabbitmqConnection()
	if connErr != nil {
		return
	}

	ch, chErr := conn.Connection.Channel()
	if chErr != nil {
		log.Fatalf("rabbitmq channel creation error : %v", chErr)
	}

	queue, queueErr := ch.QueueDeclare("notification", false, false, false, false, nil)
	if queueErr != nil {
		log.Fatalf("rabbitmq queue creation error : %v", queueErr)
	}

	msgs, msgsErr := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if msgsErr != nil {
		log.Fatalf("rabbitmq queue creation error : %v", msgsErr)
	}
	go func() {
		for msg := range msgs {
			var notification commonhelp.NotificationResp
			err := json.Unmarshal(msg.Body, &notification)
			if err != nil {
				log.Fatalf("unmarshel error : %v", err)
			}
			u.repo.StoreNotificationsOnDB(notification)
		}
	}()

	for {
		fmt.Println("Now iam working")
		time.Sleep(time.Second * 1)
	}

}
