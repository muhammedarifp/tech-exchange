package interfaces

import "github.com/muhammedarifp/content/commonHelp/response"

type ContentUserUsecase interface {
	// Post management
	CreateNewPost() (response.ContentResp, error)
	// UpdateExistingPost()
	// RemoveExistingPost()
	// LikePost()
	// PutComment()
}
