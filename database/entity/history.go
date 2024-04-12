package entity

import "gorm.io/gorm"

type Comment struct { // DB생성 로직
	gorm.Model
	Name   string
	Age    int
	Gender string
}
type Dex struct { // DB생성 로직
	gorm.Model
	Name   string
	Age    int
	Gender string
}
