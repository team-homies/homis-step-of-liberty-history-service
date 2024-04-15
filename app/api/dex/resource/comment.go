package resource

// dex
type DexEventResource struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"user_id"`
	UserName string `json:"nickname"`
	Content  string `json:"content"`
}
