package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname     string `gorm:"default:'email'; column:nickname; not null"`
	Profile      string `gorm:"column:profile; not null"`
	Provider     string `gorm:"column:provider; not null"`
	RefreshToken string `gorm:"column:refresh_token; not null"`
	IsUsed       bool   `gorm:"column:is_used; not null"`
}

type Visual struct {
	gorm.Model
	Code         string `gorm:"column:code; not null"`
	Name         string `gorm:"column:name; not null"`
	DisplayLevel uint   `gorm:"column:display_level; not null"`
	Percent      string `gorm:"column:percent; not null"`
	Description  string `gorm:"column:description; not null"`
	ImageUrl     string `gorm:"column:image_url; not null"`
}
