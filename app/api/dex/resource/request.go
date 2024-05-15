package resource

// Dex 요청값
type FindEventRequest struct {
	// 쿼리
	EventId int `json:"id" query:"id"`
}

type CreateEventRequest struct {
	// 바디
	UserId  int
	EventId int `json:"id"`
}
