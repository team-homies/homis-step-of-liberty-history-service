package resource

type FindAllCommentRequest struct {
	// 라우터에 포함될 EventId Param
	EventId uint64 `json:"event_id" param:"id"`
}

type CreateCommentRequest struct {
	// 라우터에 포함될 EventId Param
	EventId uint64 `json:"event_id" param:"id"`
	// 등록할 UserId와 Content body
	UserId  uint64 `json:"user_id"`
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	Id uint64 `json:"id"`
	// 라우터에 포함될 Id Param
	EventId uint64 `json:"event_id" param:"id"`
	// 수정할 UserId와 Content body
	UserId  uint64 `json:"user_id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	Id uint64 `json:"id"`
	// 라우터에 포함될 Id Param
	EventId uint64 `json:"event_id" param:"id"`
	// 삭제할 UserId body
	UserId uint64 `json:"user_id"`
}
