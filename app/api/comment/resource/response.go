package resource

// comment 반환값
type FindAllCommentResponse struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
