package interfaces

import (
	"context"

	"github.com/muhammedarifp/content/commonHelp/response"
)

type ContentUserRepository interface {
	CreateNewPost(ctx context.Context) (response.ContentResp, error)
}
