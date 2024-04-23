package resource

// dex 반환값
type GetDexEventResponse struct {
	Dex DexEventResource `json:"dex"`
}

type CreateDexEventResponse struct {
	Dex DexEventResource `json:"dex"`
}
