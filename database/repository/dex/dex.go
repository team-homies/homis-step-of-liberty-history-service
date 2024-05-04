package dex

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexRepository interface {
	FindDexEventById(id int64) (res *entity.Event, err error)
	FindDexDetailById(id int64) (res *entity.Detail, err error)
	FindUserDexById(dexId int64, userId int64) (res int64, err error)
	CreateUserDexById(dexId int64, userId int64) (err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

func (g *gormDexRepository) FindDexEventById(id int64) (res *entity.Event, err error) {
	// 1. 쿼리작성
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.dex_id;

	// 2. gorm로직
	tx := g.db
	tx.Model(&entity.Event{}).Select("name", "level").Where("id = ?", id).First(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexRepository) FindDexDetailById(id int64) (res *entity.Detail, err error) {
	// 1. 쿼리작성
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.dex_id;

	// 2. gorm로직
	tx := g.db
	tx.Model(&entity.Detail{}).Select("define", "outline", "place", "background", "image_url").Where("dex_id = ?", id).First(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormDexRepository) FindUserDexById(dexId int64, userId int64) (res int64, err error) {
	// 1. 쿼리작성
	// select * from userdex where dex_id = 1 and user_id = 1

	// 2. gorm로직
	var dexCount int64
	tx := g.db
	tx.Model(&entity.UserDex{}).Where("dex_id = ?", dexId).Where("user_id = ?", userId).Count(&dexCount)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return dexCount, nil
}

func (g *gormDexRepository) CreateUserDexById(dexId int64, userId int64) (err error) {
	// 1. 쿼리작성
	// insert into userdex (dex_id, user_id)
	// values (1, 1)

	// 2. gorm로직
	tx := g.db.Begin()
	tx.Model(&entity.UserDex{}).Create(&entity.UserDex{
		EventId: int(dexId),
		UserId:  int(userId),
	})
	if tx.Error != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
