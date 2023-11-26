package requests

type CreateNewPostRequest struct {
	ThumbnailImg     string
	Title            string
	Body             string
	Is_showReactions bool
	Is_premium       bool
}
