package resource

type GetQuoteResponse struct {
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}
type GetTagsResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

	// type GetTagsResponse struct {
	// 	Tags []TagResource `json:"tags"`
	// }
	
