package resource

// comment
type CommentResource struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	UserName string `json:"nickname"`
	Content  string `json:"content"`
}
