package resource

type DexEventResponse struct {
	EventId     int64  `json:"event_id"`
	Name        string `json:"name"`
	Level       string `json:"level"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	IsUsed      bool   `json:"is_used"`
	EventDetail Detail
}

type Detail struct {
	DetailId   string `json:"detail_id"`
	Define     string `json:"define"`
	Outline    string `json:"outline"`
	Place      string `json:"place"`
	Background string `json:"background"`
	ImageUrl   string `json:"image_url"`
}
