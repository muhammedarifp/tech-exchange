package interfaces

import "github.com/muhammedarifp/content/domain"

type AdminContentUseCase interface {
	GetallPosts(page int) ([]domain.Contents, error)
	RemovePost(postid, userid string) (domain.Contents, error)
	CreateNewTag(tsg string) (domain.Tags, error)
}
