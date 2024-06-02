package resource

type CommentResource struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
