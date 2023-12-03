package usecases

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/rabbitmq"
	"github.com/muhammedarifp/content/repository/interfaces"
	services "github.com/muhammedarifp/content/usecases/interfaces"
	"github.com/rabbitmq/amqp091-go"
)

type ContentUserUsecase struct {
	userRepo interfaces.ContentUserRepository
}

func NewContentUserUsecase(repo interfaces.ContentUserRepository) services.ContentUserUsecase {
	return &ContentUserUsecase{userRepo: repo}
}

func (c *ContentUserUsecase) CreatePost(userid string, post requests.CreateNewPostRequest) (domain.Contents, error) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	val, err := c.userRepo.CreateNewPost(ctx, userid, post)
	if err != nil {
		return val, err
	}

	return val, nil
}

func (c *ContentUserUsecase) CreateComment(userid, postid, text string) (domain.Contents, error) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	content, err := c.userRepo.CreateComment(ctx, postid, userid, text)
	if err != nil {
		return content, err
	}

	return content, nil
}

func (c *ContentUserUsecase) LikePost(postid string) (domain.Contents, error) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	content, err := c.userRepo.LikePost(ctx, postid)

	// Send notification
	go func() {
		conn, _ := rabbitmq.NewRabbitmqConnection()

		ch, cherr := conn.Channel()
		if cherr != nil {
			log.Fatalf("channel creation error : %v", cherr)
		}

		queue, queueErr := ch.QueueDeclare("notification", false, false, false, false, nil)
		if queueErr != nil {
			log.Fatalf("channel creation error : %v", queueErr)
		}

		notification, marshelErr := json.Marshal(requests.NotificationReq{
			UserID:      uint(10),
			Title:       "New Like ... ",
			Body:        "New Like on your post ...",
			IsImportent: true,
		})

		if marshelErr != nil {
			log.Fatalf("marshel error : %v", marshelErr)
		}

		if publishErr := ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        notification,
		}); publishErr != nil {
			log.Fatalf("publish error : %v", publishErr)
		}
	}()

	return content, err
}
