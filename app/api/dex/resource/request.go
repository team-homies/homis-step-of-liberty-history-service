package resource

// Dex 요청값
type FindDexEventRequest struct {
	// 쿼리
	EventId int `json:"id" query:"id"`
}

type CreateDexEventRequest struct {
	// 바디
	UserId  int
	EventId int `json:"id"`
}
