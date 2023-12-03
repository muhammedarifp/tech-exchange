package commonhelp

type NotificationResp struct {
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsImportent bool   `json:"is_importent"`
}