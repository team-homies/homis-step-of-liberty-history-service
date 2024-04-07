package entity

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name    string `gorm:"column:name; not null"`
	Level   string `gorm:"column:level; not null"`
	IsUsed  bool   `gorm:"column:is_used; not null"`
	Detail  Detail
	Userdex Userdex
	Mapping Mapping
	Comment Comment
}

func (Event) TableName() string {
	return "event"
}

type Detail struct {
	gorm.Model
	DexId      int    `gorm:"column:dex_id; foreignKey:EventId; not null"`
	Define     string `gorm:"column:define"`
	Outline    string `gorm:"column:outline"`
	Place      string `gorm:"column:place"`
	Background string `gorm:"column:background"`
	// Passing    string `gorm:"type:varchar"`
	// Outcome    string `gorm:"type:varchar"`
	// Assessment string `gorm:"type:varchar"`
	ImageUrl string `gorm:"column:image_url"`
}

func (Detail) TableName() string {
	return "detail"
}

type Userdex struct {
	gorm.Model
	DexId  int `gorm:"column:dex_id; foreignKey:Id; not null"`
	UserId int `gorm:"column:user_id; not null"`
}

func (Userdex) TableName() string {
	return "userdex"
}

type Tag struct {
	gorm.Model
	Name    string `gorm:"column:name; not null"`
	Mapping Mapping
}

func (Tag) TableName() string {
	return "tag"
}

type Mapping struct {
	Id    int `gorm:"primaryKey; column:id"`
	DexId int `gorm:"column:dex_id; foreignKey:EventId; not null"`
	TagId int `gorm:"column:tag_id; foreignKey:TagId; not null"`
}

func (Mapping) TableName() string {
	return "mapping"
}

type Comment struct {
	gorm.Model
	DexId   int    `gorm:"column:dex_id; foreignKey:EventId; not null"`
	UserId  int    `gorm:"column:user_id; not null"`
	Content string `gorm:"column:content; not null"`
}

func (Comment) TableName() string {
	return "comment"
}

type Quote struct {
	gorm.Model
	Content  string `gorm:"column:content; not null"`
	ImageUrl string `gorm:"column:image_url; not null"`
	StartAt  string `gorm:"column:start_at; not null"`
	EndAt    string `gorm:"column:end_at; not null"`
}

func (Quote) TableName() string {
	return "quote"
}
