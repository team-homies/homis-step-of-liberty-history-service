package resource

// dex 반환값
type FindEventResponse struct {
	Name   string `json:"name"`
	Level  string `json:"level"`
	Detail FindDetailResponse
}
type FindDetailResponse struct {
	Define     string `json:"define"`
	Outline    string `json:"outline"`
	Place      string `json:"place"`
	Background string `json:"background"`
	ImageUrl   string `json:"image_url"`
}

type CreateEventResponse struct {
	Dex EventResource `json:"dex"`
}
