package entity

import "gorm.io/gorm"

type Patient struct { // DB생성 로직
	gorm.Model
	Name   string
	Age    int
	Gender string
}
