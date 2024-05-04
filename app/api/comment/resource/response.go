package resource

// comment 반환값
type GetAllCommentResponse struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type CreateCommentResponse struct {
	Comment CommentResource `json:"comment"`
}

type UpdateCommentResponse struct {
	Comment CommentResource `json:"comment"`
}

type DeleteCommentResponse struct {
	Comment CommentResource `json:"comment"`
}
