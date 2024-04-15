package resource

// dex 반환값
type GetDexResponse struct {
	Dex DexEventResource `json:"dex"`
}

type CreateDexResponse struct {
	Dex DexEventResource `json:"dex"`
}
