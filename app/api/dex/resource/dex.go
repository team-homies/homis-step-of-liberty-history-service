package resource

import "gorm.io/gorm"

// dex
type DexEventResource struct {
	gorm.Model
	Name   string `json:"name"`
	Level  string `json:"level"`
	IsUsed bool   `json:"is_used"`
	Detail DetailEventResource
}

type DetailEventResource struct {
	gorm.Model
	EventId    int    `json:"dex_id"`
	Define     string `json:"define"`
	Outline    string `json:"outline"`
	Place      string `json:"place"`
	Background string `json:"background"`
	ImageUrl   string `json:"image_url"`
}
