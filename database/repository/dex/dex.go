package dex

import (
	"main/app/api/dex/resource"
	"main/database/entity"

	"gorm.io/gorm"
)

// Dex 레포지토리 인터페이스
type DexRepository interface {
	FindDexEventByEventId(eventId int) (res *resource.EventJoinResource, err error)
	FindUserDexByEventId(eventId int, userId int) (res int, err error)
	CreateUserDexByEventId(eventId int, userId int) (err error)
	GetQuote() (quote []entity.Quote, err error)
	GetTags() (result []entity.Tag, err error)
	CountEvents() (count int64, err error)
	CountUserEvents(userId uint64) (count int64, err error)
}

type gormDexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) DexRepository {
	return &gormDexRepository{db}
}

// [사건 내용 조회] 사건 id로 조회 select문
func (g *gormDexRepository) FindDexEventByEventId(eventId int) (res *resource.EventJoinResource, err error) {
	// select  e.name, e.level, d.define, d.outline, d.place, d.background, d.image_url
	//   from "event" e, "detail" d
	//  where e.id = d.event_id;

	// 1. gorm로직
	err = g.db.Model(&entity.Event{}).
		Select("Event.name, Event.level, Detail.define, Detail.outline, Detail.place, Detail.background, Detail.image_url").
		Joins("JOIN Detail ON Detail.event_id = Event.id AND Event.id = ?", eventId).
		Find(&res).Error

	if err != nil {
		return
	}
	return
}

// [사용자 사건 수집 등록] 사건 id로 조회 select문 : 사건보유 여부 위함
func (g *gormDexRepository) FindUserDexByEventId(eventId int, userId int) (res int, err error) {
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
func (g *gormDexRepository) CreateUserDexByEventId(eventId int, userId int) (err error) {
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

// 태그 목록 조회
func (g *gormDexRepository) GetTags() (result []entity.Tag, err error) {
	// select "id", name
	//   from tag t;

	tx := g.db.Select("id", "name").Find(&result)

	return result, tx.Error

}

// 전체 사건의 개수
func (g *gormDexRepository) CountEvents() (count int64, err error) {
	// 	select count(*)
	//    from "event" e
	//   where is_used = true ;
	err = g.db.Model(&entity.Event{}).Where("is_used = true").Count(&count).Error

	return
}

// 유저 사건의 전체 개수
func (g *gormDexRepository) CountUserEvents(userId uint64) (count int64, err error) {
	// select count(*)
	//   from userdex u
	//  where user_id = 1
	// 	  and deleted_at is null ;
	err = g.db.Model(&entity.UserDex{}).Where("user_id = ? AND deleted_at is null", userId).Count(&count).Error

	return
}

// 명언 전체 조회
func (g *gormDexRepository) GetQuote() (quote []entity.Quote, err error) {
	// 	select *
	//    from "quote" q
	err = g.db.Find(&quote).Error

	return
}
