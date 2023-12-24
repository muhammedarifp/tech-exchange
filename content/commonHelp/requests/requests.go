package requests

type CreateNewPostRequest struct {
	Title            string   `json:"title"`
	Body             string   `json:"body"`
	Is_showReactions bool     `json:"is_show_reactions"`
	Is_premium       bool     `json:"is_premium"`
	Labels           []string `json:"labels"`
}

type UpdatePostRequest struct {
	Postid           string `json:"postid"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Is_showReactions bool   `json:"is_show_reactions"`
	Is_premium       bool   `json:"is_premium"`
}

type FollowTagReq struct {
	Tags []string `json:"tags"`
}
