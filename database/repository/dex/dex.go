package dex

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexRepository interface {
	FindDexEventById(id int) (res *entity.Event, err error)
	FindDexDetailById(id int) (res *entity.Detail, err error)
	FindUserDexById(eventId int, userId int) (res int, err error)
	CreateUserDexById(eventId int, userId int) (err error)
	GetTags() (result []entity.Tag, err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

// [사건 내용 조회] 사건 id로 조회 select문
func (g *gormDexRepository) FindDexEventById(id int) (res *entity.Event, err error) {
	// 1. 쿼리작성
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.event_id;

	// 2. gorm로직
	tx := g.db
	err = tx.Model(&entity.Event{}).Select("name", "level").Where("id = ?", id).First(&res).Error

	if err != nil {
		return
	}
	return
}

// [사건 내용 조회] 사건 detail id로 조회 select문
func (g *gormDexRepository) FindDexDetailById(id int) (res *entity.Detail, err error) {
	// 1. 쿼리작성
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.event_id;

	// 2. gorm로직
	tx := g.db
	err = tx.Model(&entity.Detail{}).Select("define", "outline", "place", "background", "image_url").Where("event_id = ?", id).First(&res).Error

	if err != nil {
		return
	}
	return
}

// [사용자 사건 수집 등록] 사건 id로 조회 select문 : 사건보유 여부 위함
func (g *gormDexRepository) FindUserDexById(eventId int, userId int) (res int, err error) {
	// 1. 쿼리작성
	// select * from userdex where event_id = 1 and user_id = 1

	// 2. gorm로직
	var dexCount int64
	tx := g.db
	err = tx.Model(&entity.UserDex{}).Where("event_id = ?", eventId).Where("user_id = ?", userId).Count(&dexCount).Error
	if err != nil {
		return
	}
	return
}

// [사용자 사건 수집 등록] 사건 id로 등록 post문
func (g *gormDexRepository) CreateUserDexById(eventId int, userId int) (err error) {
	// 1. 쿼리작성
	// insert into userdex (event_id, user_id)
	// values (1, 1)

	// 2. gorm로직
	tx := g.db.Begin()
	err = tx.Model(&entity.UserDex{}).Create(&entity.UserDex{
		EventId: eventId,
		UserId:  userId,
	}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (g *gormDexRepository) GetTags() (result []entity.Tag, err error) {
	// select "id", name
	// from tag t;

	tx := g.db.Select("id", "name").Find(&result)

	return result, tx.Error

}
