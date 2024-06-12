package resource

// dex 반환값
type FindEventResponse struct {
	Name   string `json:"name"`
	Level  string `json:"level"`
	Detail FindDetailResponse
}
type FindDetailResponse struct {
	Place        string `json:"place"`
	Situation    string `json:"situation"`
	Organization string `json:"organization"`
	Person       string `json:"person"`
	Content      string `json:"content"`
	Appraisal    string `json:"appraisal"`
	Reference    string `json:"reference"`
	ImageUrl     string `json:"image_url"`
}
type CreateEventResponse struct {
	Dex EventResource `json:"dex"`
}
type GetQuoteResponse struct {
	Id       uint   `json:"id"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
}
type GetTagsResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
