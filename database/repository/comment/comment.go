package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Comment 레포지토리 인터페이스
type CommentRepository interface {
	FindAll(Id int) (result []entity.Comment, err error)
	Create(comment *entity.Comment) (err error)
	Update(comment *entity.Comment) (err error)
	Delete(comment *entity.Comment) (err error)
}

type gormCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &gormCommentRepository{db}
}

// [혈서 목록 조회] Param : Id
func (g *gormCommentRepository) FindAll(Id int) (result []entity.Comment, err error) {
	// 1. 쿼리 작성
	// select (id, event_id, user_id, content) from comment

	// 2. gorm 적용
	tx := g.db
	err = tx.Model(&entity.Comment{}).Select("user_id", "content").Where("id = ?", Id).Find(&result).Error
	if err != nil {
		return
	}
	return
}

// [혈서 등록] Param : Id, body : UserId, Content
func (g *gormCommentRepository) Create(comment *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// insert into comment (id, event_id, user_id, content)
	// values (1, 1, 1, "내용")

	// 2. gorm 적용
	tx := g.db.Begin()
	err = tx.Model(&entity.Comment{}).Create(&comment).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// [혈서 수정] Param : Id, body : UserId, Content
func (g *gormCommentRepository) Update(comment *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// update "comment"
	//    set content = '수정내용'
	//  where content is not null;

	// 2. gorm 적용
	tx := g.db.Begin()
	err = tx.Model(&entity.Comment{}).Where("id = ?", comment.ID).Updates(&comment).Error

	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()

	return
}

// [혈서 삭제] Param : Id, body : UserId
func (g *gormCommentRepository) Delete(comment *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// 	delete
	// 	from "comment"
	//  where id = 7;

	// deleted_at에 값이 null인 값을 삭제
	var result entity.Comment
	// 2. gorm 적용
	tx := g.db.Begin()
	err = tx.Model(&entity.Comment{}).Where("id = ?", comment.ID).Delete(&result).Error

	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
