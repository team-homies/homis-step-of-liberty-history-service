package resource

// dex
type DexResource struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"user_id"`
	UserName string `json:"nickname"`
	Content  string `json:"content"`
}
