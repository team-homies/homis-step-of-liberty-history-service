package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Comment 레포지토리 인터페이스
type CommentRepository interface {
	FindAll(eventId int) (res []entity.Comment, err error)
	Create(eventId int, req *entity.Comment) (err error)
	Update(eventId int, req *entity.Comment) (err error)
	Delete(eventId int, req *entity.Comment) (res *entity.Comment, err error)
}

type gormCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &gormCommentRepository{db}
}

// [혈서 목록 조회] query : EventId
func (g *gormCommentRepository) FindAll(eventId int) (res []entity.Comment, err error) {
	// 1. 쿼리 작성
	// select (id, event_id, user_id, content) from comment

	// 2. gorm 적용
	tx := g.db
	tx.Model(&entity.Comment{}).Select("user_id", "content").Where("event_id = ?", eventId).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

// [혈서 등록] query : EventId, body : UserId, Content
func (g *gormCommentRepository) Create(eventId int, req *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// insert into comment (id, event_id, user_id, content)
	// values (1, 1, 1, "내용")

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Create(&entity.Comment{
		UserId:  req.UserId,
		Content: req.Content,
	})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return tx.Error
}

// [혈서 수정] query : EventId, body : UserId, Content
func (g *gormCommentRepository) Update(eventId int, req *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// update "comment"
	//    set content = '수정내용'
	//  where content is not null;

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Where("event_id = ?", req.EventId).Updates(req.Content)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()

	return err
}

// [혈서 삭제] query : EventId, body : UserId
func (g *gormCommentRepository) Delete(eventId int, req *entity.Comment) (res *entity.Comment, err error) {
	// 1. 쿼리 작성
	// 	delete
	// 	from "comment"
	//  where event_id = 7 and user_id = 1;

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Where("event_id = ?", eventId).Where("user_id = ?", req.UserId).First(&res)
	tx.Delete(&res)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return res, nil
}
