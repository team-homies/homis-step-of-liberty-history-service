package resource

type GetQuoteResponse struct {
	Id       uint   `json:"id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}
type GetTagsResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}