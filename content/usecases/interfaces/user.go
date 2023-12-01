package interfaces

import (
	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/domain"
)

type ContentUserUsecase interface {
	// Post management
	CreatePost(userid string, post requests.CreateNewPostRequest) (domain.Contents, error)
	CreateComment(userid, postid, text string) (domain.Contents, error)
	// UpdateExistingPost()
	// RemoveExistingPost()
	// LikePost()
	// PutComment()
}
