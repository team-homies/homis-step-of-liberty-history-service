package resource

// comment 반환값
type FindAllCommentResponse struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type CreateCommentResponse struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type UpdateCommentResponse struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type DeleteCommentResponse struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
