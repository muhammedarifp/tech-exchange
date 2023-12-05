package interfaces

import "github.com/muhammedarifp/content/domain"

type AdminContentUseCase interface {
	GetallPosts(page int) ([]domain.Contents, error)
}
