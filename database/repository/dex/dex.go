package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexRepository interface {
	Get() (res []entity.Dex, err error)
	Create(req *entity.Dex) (err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

func (g *gormDexRepository) Get() (res []entity.Dex, err error) {
	// panic("error!!")
	tx := g.db
	tx.Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexRepository) Create(req *entity.Dex) (err error) {
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
