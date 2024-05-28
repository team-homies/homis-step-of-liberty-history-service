package resource

type FindAllCommentRequest struct {
	// 라우터에 포함될 EventId Param
	EventId int `json:"event_id" param:"id"`
}

type CreateCommentRequest struct {
	// 라우터에 포함될 EventId Param
	EventId int `json:"event_id" param:"id"`
	// 등록할 UserId와 Content body
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type UpdateCommentRequest struct {
	// 라우터에 포함될 Id Param
	Id int `json:"id" param:"id"`
	// 수정할 UserId와 Content body
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

type DeleteCommentRequest struct {
	// 라우터에 포함될 Id Param
	Id int `json:"id" param:"id"`
	// 삭제할 UserId body
	UserId int `json:"user_id"`
}
