package interfaces

import (
	"github.com/muhammedarifp/content/commonHelp/requests"
	"github.com/muhammedarifp/content/domain"
)

type ContentUserUsecase interface {
	// Post management
	CreatePost(userid string, post requests.CreateNewPostRequest) (domain.Contents, error)
	CreateComment(userid, postid, text string) (domain.Contents, error)
	LikePost(postid string, user_id string) (domain.Contents, error)
	UpdatePost(post requests.UpdatePostRequest, userid string) (domain.Contents, error)
	RemovePost(postid, userid string) (domain.Contents, error)
	GetUserPosts(userid string, page int) ([]domain.Contents, error)
	GetallPosts(page int) ([]domain.Contents, error)
}
