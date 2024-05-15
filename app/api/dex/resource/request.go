package resource

// Dex 요청값
type GetDexEventRequest struct {
	// 쿼리
	Id int `json:"id" query:"id"`
}

type CreateDexEventRequest struct {
	// 바디
	UserId  int
	EventId int `json:"id"`
}
