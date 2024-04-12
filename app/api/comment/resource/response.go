package resource

// comment 반환값
type GetAllCommentResponse struct {
	Comment []CommentResource `json:"comments"`
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
