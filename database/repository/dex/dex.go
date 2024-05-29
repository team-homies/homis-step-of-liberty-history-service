package dex

import (
	"main/database/entity"

	"gorm.io/gorm"
)

type DexRepository interface {
	FindQuote(Id uint) (quote *entity.Quote, err error)
	GetTags() (result []entity.Tag, err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

func (g *gormDexRepository) FindQuote(Id uint) (quote *entity.Quote, err error) {
	// 	select *
	//    from "quote" q
	//   where id = 1;
	tx := g.db.Where("id = ?", Id).Find(&quote)

	return quote, tx.Error
}
func (g *gormDexRepository) GetTags() (result []entity.Tag, err error) {
	// select "id", name
	// from tag t;

	tx := g.db.Select("id", "name").Find(&result)

	return result, tx.Error

}
