package resource

// comment 요청값
type GetAllCommentRequest struct {
	Id int `json:"id" query:"id"`
}

type CreateCommentRequest struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	Id      int    `json:"id" query:"id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	Id int `json:"id" query:"id"`
}
