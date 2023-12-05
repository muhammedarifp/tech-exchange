package usecases

import (
	"context"
	"errors"
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
