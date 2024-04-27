package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexEventRepository interface {
	Get() (res []entity.Event, err error)
	Create(req *entity.Event) (err error)
}

type gormDexEventRepository struct {
	db *gorm.DB
}

func NewDexEventRepository(db *gorm.DB) DexEventRepository {
	return &gormDexEventRepository{db}
}

func (g *gormDexEventRepository) Get() (res []entity.Event, err error) {
	// panic("error!!")
	tx := g.db
	tx.Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexEventRepository) Create(req *entity.Event) (err error) {
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
