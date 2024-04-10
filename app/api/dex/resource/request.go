package resource

// Dex 요청값
type GetDexRequest struct {
	Id uint64 `json:"id" query:"id"`
}

type GetDexsRequest struct {
	Id uint64 `json:"id" query:"id"`
}

type CreateDexRequest struct {
	Id      uint64 `json:"id" query:"id"`
	Content string `json:"content"`
}

type UpdateDexRequest struct {
	Id      uint64 `json:"id" query:"id"`
	Content string `json:"content"`
}

type DeleteDexRequest struct {
	Id uint64 `json:"id" query:"id"`
}
