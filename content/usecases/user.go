package usecases

import (
	"context"
	"time"

	"github.com/muhammedarifp/content/commonHelp/response"
	"github.com/muhammedarifp/content/repository/interfaces"
	services "github.com/muhammedarifp/content/usecases/interfaces"
)

type ContentUserUsecase struct {
	userRepo interfaces.ContentUserRepository
}

func NewContentUserUsecase(repo interfaces.ContentUserRepository) services.ContentUserUsecase {
	return &ContentUserUsecase{userRepo: repo}
}

func (c *ContentUserUsecase) CreateNewPost() (response.ContentResp, error) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	val, err := c.userRepo.CreateNewPost(ctx)
	if err != nil {
		return val, err
	}

	return val, nil
}
