package entity

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name    string `gorm:"column:name;not null"`
	Level   string `gorm:"column:level;not null"`
	IsUsed  bool   `gorm:"column:is_used;not null"`
	Detail  Detail
	Userdex UserDex
	Mapping Mapping
	Comment Comment
}

func (Event) TableName() string {
	return "event"
}

type Detail struct {
	gorm.Model
	EventId    int    `gorm:"column:dex_id;foreignKey:EventId;"`
	Define     string `gorm:"column:define"`
	Outline    string `gorm:"column:outline"`
	Place      string `gorm:"column:place"`
	Background string `gorm:"column:background"`
	ImageUrl   string `gorm:"column:image_url"`
}

func (Detail) TableName() string {
	return "detail"
}

type UserDex struct {
	gorm.Model
	EventId int `gorm:"column:dex_id; foreignKey:EventId;"`
	UserId  int `gorm:"column:user_id;not null"`
}

func (UserDex) TableName() string {
	return "userdex"
}

type Tag struct {
	gorm.Model
	Name    string `gorm:"column:name;not null"`
	Mapping Mapping
}

func (Tag) TableName() string {
	return "tag"
}

type Mapping struct {
	Id      int `gorm:"primaryKey;column:id"`
	EventId int `gorm:"column:dex_id;foreignKey:EventId;"`
	TagId   int `gorm:"column:tag_id;foreignKey:TagId;"`
}

func (Mapping) TableName() string {
	return "mapping"
}

type Comment struct {
	gorm.Model
	EventId int    `gorm:"column:dex_id;foreignKey:EventId;"`
	UserId  int    `gorm:"column:user_id;not null"`
	Content string `gorm:"column:content;not null"`
}

func (Comment) TableName() string {
	return "comment"
}

type Quote struct {
	gorm.Model
	Content  string `gorm:"column:content;not null"`
	ImageUrl string `gorm:"column:image_url;not null"`
}

func (Quote) TableName() string {
	return "quote"
}
