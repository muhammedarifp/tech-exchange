package usecases

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/muhammedarifp/content/domain"
	"github.com/muhammedarifp/content/repository/interfaces"
	services "github.com/muhammedarifp/content/usecases/interfaces"
)

type AdminContentUsecase struct {
	repo interfaces.AdminContentRepo
}

func NewAdminContentUsecase(repo interfaces.AdminContentRepo) services.AdminContentUseCase {
	return &AdminContentUsecase{
		repo: repo,
	}
}

func (u *AdminContentUsecase) GetallPosts(page int) ([]domain.Contents, error) {
	if page <= 0 {
		return []domain.Contents{}, errors.New("Invalid page number")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	posts, repoErr := u.repo.GetallPosts(ctx, page)
	if repoErr != nil {
		return []domain.Contents{}, repoErr
	}

	return posts, nil
}

func (c *AdminContentUsecase) RemovePost(postid, userid string) (domain.Contents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	content, repoErr := c.repo.RemovePost(ctx, postid, userid)
	if repoErr != nil {
		return content, repoErr
	}

	return content, repoErr
}

func (c *AdminContentUsecase) CreateNewTag(tag string) (domain.Tags, error) {
	tag = strings.ToLower(tag)
	if tag == "" {
		return domain.Tags{}, errors.New("input value is incorrect")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	newTag, err := c.repo.AddNewTag(ctx, tag)
	if err != nil {
		return domain.Tags{}, err
	}

	return newTag, nil
}
