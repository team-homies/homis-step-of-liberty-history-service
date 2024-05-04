package dex

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexRepository interface {
	FindDexEventById(id int64) (res *entity.Event, err error)
	FindDexDetailById(id int64) (res *entity.Detail, err error)
	Create(req *entity.Event) (err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

func (g *gormDexRepository) FindDexEventById(id int64) (res *entity.Event, err error) {
	// panic("error!!")
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.dex_id;
	tx := g.db
	tx.Model(&entity.Event{}).Select("name", "level").Where("id = ?", id).First(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexRepository) FindDexDetailById(id int64) (res *entity.Detail, err error) {
	// panic("error!!")
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.dex_id;
	tx := g.db
	tx.Model(&entity.Detail{}).Select("define", "outline", "place", "background", "image_url").Where("dex_id = ?", id).First(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexRepository) Create(req *entity.Event) (err error) {
	// panic("")
	tx := g.db.Begin()
	tx.Create(&req)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return tx.Error
}
