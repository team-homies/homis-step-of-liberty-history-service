package resource

// dex 반환값
type GetDexResponse struct {
	Dex DexResource `json:"dex"`
}

type GetDexsResponse struct {
	Dex []DexResource `json:"dexs"`
}

type CreateDexResponse struct {
	Dex DexResource `json:"dex"`
}

type UpdateDexResponse struct {
	Dex DexResource `json:"dex"`
}

type DeleteDexResponse struct {
	Dex DexResource `json:"dex"`
}
