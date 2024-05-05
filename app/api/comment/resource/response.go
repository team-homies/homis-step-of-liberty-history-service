package resource

// comment 반환값
type GetAllCommentResponse struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type CreateCommentResponse struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type UpdateCommentResponse struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type DeleteCommentResponse struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
