package interfaces

import (
	"context"

	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/domain"
)

type ContentUserRepository interface {
	CreateNewPost(ctx context.Context, userid string, post requests.CreateNewPostRequest) (domain.Contents, error)
	CreateComment(ctx context.Context, post_id, userid string, text string) (domain.Contents, error)
	LikePost(ctx context.Context, postid string) (domain.Contents, error)
}
