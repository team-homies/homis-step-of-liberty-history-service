package resource

// comment
type CommentResource struct {
	EventId int    `json:"event_id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
