package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
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
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	val, err := c.userRepo.CreateNewPost(ctx, userid, post)
	if err != nil {
		return val, err
	}

	return val, nil
}

func (c *ContentUserUsecase) CreateComment(userid, postid, text string) (domain.Contents, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	content, err := c.userRepo.CreateComment(ctx, postid, userid, text)
	if err != nil {
		return content, err
	}

	return content, nil
}

func (c *ContentUserUsecase) LikePost(postid, user_id string) (domain.Contents, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	content, respoErr := c.userRepo.LikePost(ctx, postid, user_id)
	var emptyContent domain.Contents

	if respoErr != nil || content.UserID == 0 {
		return emptyContent, errors.New("Internal server error try again later")
	}

	useridUInt, strconvErr := strconv.Atoi(user_id)
	if strconvErr != nil {
		return emptyContent, errors.New("userid is not convertable")
	}

	if content.UserID != uint(useridUInt) {
		// Send notification
		go func() {
			conn, connErr := rabbitmq.NewRabbitmqConnection()
			if connErr != nil {
				log.Printf("rabbitmq connection error: %v", connErr)
				return
			}

			ch, cherr := conn.Channel()
			if cherr != nil {
				log.Printf("rabbitmq connection error: %v", ch)
				return
			}

			queue, queueErr := ch.QueueDeclare("notification", false, false, false, false, nil)
			if queueErr != nil {
				log.Printf("queue creation error: %v", queueErr)
				return
			}

			notification, marshelErr := json.Marshal(requests.NotificationReq{
				UserID:      content.UserID,
				LikedUserID: uint(useridUInt),
				Title:       "New Like Notification",
				Body:        "Your post received a new like!",
				IsImportent: true,
			})

			if marshelErr != nil {
				log.Printf("marshal error: %v", marshelErr)
				return
			}

			if publishErr := ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp091.Publishing{
				ContentType: "application/json",
				Body:        notification,
			}); publishErr != nil {
				log.Printf("publish error: %v", publishErr)
				return
			}
		}()
	}

	return content, nil
}

func (c *ContentUserUsecase) UpdatePost(post requests.UpdatePostRequest, userid string) (domain.Contents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	content, repoErr := c.userRepo.UpdatePost(ctx, post, userid)
	if repoErr != nil {
		return content, repoErr
	}

	return content, repoErr
}

func (c *ContentUserUsecase) RemovePost(postid, userid string) (domain.Contents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	content, repoErr := c.userRepo.RemovePost(ctx, postid, userid)
	if repoErr != nil {
		return content, repoErr
	}

	return content, repoErr
}

func (c *ContentUserUsecase) GetUserPosts(userid string, page int) ([]domain.Contents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	contents, repoErr := c.userRepo.GetUserPosts(ctx, userid, page)
	if repoErr != nil {
		fmt.Println("repo error found")
		return contents, repoErr
	}

	return contents, repoErr
}

func (u *ContentUserUsecase) GetallPosts(page int) ([]domain.Contents, error) {
	if page <= 0 {
		return []domain.Contents{}, errors.New("Invalid page number")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	posts, repoErr := u.userRepo.GetallPosts(ctx, page)
	if repoErr != nil {
		return []domain.Contents{}, repoErr
	}

	return posts, nil
}
