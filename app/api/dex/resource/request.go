package resource

// Dex 요청값
type FindEventRequest struct {
	// 패스파라미터
	EventId int `json:"id" param:"id"`
}

type CreateEventRequest struct {
	// 바디
	UserId  int
	EventId int `json:"id"`
}

type GetQuoteRequest struct {
	Id uint `json:"id"`
}
