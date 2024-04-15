package resource

// Dex 요청값
type GetDexEventRequest struct {
	Id uint64 `json:"id" query:"id"`
}

type CreateDexEventRequest struct {
	Id      uint64 `json:"id" query:"id"`
	Content string `json:"content"`
}
