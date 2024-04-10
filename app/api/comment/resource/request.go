package resource

// comment 요청값
type GetCommentRequest struct {
	Id uint64 `json:"id" query:"id"`
}

type GetCommentsRequest struct {
	Id uint64 `json:"id" query:"id"`
}

type CreateCommentRequest struct {
	Id      uint64 `json:"id" query:"id"`
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	Id      uint64 `json:"id" query:"id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	Id uint64 `json:"id" query:"id"`
}
