package usecases

import (
	"context"
	"time"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	services "github.com/muhammedarifp/content/usecases/interfaces"
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
