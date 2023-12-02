package commonhelp

type NotificationResp struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsImportent bool   `json:"is_importent"`
}
