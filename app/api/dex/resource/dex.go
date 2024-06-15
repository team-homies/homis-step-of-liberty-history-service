package resource

import "gorm.io/gorm"

// dex
type EventResource struct {
	gorm.Model
	Name   string `json:"name"`
	Level  string `json:"level"`
	IsUsed bool   `json:"is_used"`
	Detail DetailResource
}

type DetailResource struct {
	gorm.Model
	EventId    int    `json:"event_id"`
	Define     string `json:"define"`
	Outline    string `json:"outline"`
	Place      string `json:"place"`
	Background string `json:"background"`
	ImageUrl   string `json:"image_url"`
}

type EventJoinResource struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Define     string `json:"define"`
	Outline    string `json:"outline"`
	Place      string `json:"place"`
	Background string `json:"background"`
	ImageUrl   string `json:"image_url"`
}
