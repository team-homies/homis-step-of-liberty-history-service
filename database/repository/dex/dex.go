package dex

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type DexRepository interface {
	GetQuote() (quote []entity.Quote, err error)
	GetTags() (result []entity.Tag, err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

// 태그 목록 조회
func (g *gormDexRepository) GetTags() (result []entity.Tag, err error) {
	// select "id", name
	//   from tag t;

	tx := g.db.Select("id", "name").Find(&result)

	return result, tx.Error

}

// 명언 전체 조회
func (g *gormDexRepository) GetQuote() (quote []entity.Quote, err error) {
	// 	select *
	//    from "quote" q
	err = g.db.Find(&quote).Error

	return
}
