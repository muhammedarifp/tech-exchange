package interfaces

import (
	"context"

	"github.com/muhammedarifp/content/domain"
)

type AdminContentRepo interface {
	GetallPosts(ctx context.Context, page int) ([]domain.Contents, error)
	RemovePost(ctx context.Context, postid, userid string) (domain.Contents, error)
}
