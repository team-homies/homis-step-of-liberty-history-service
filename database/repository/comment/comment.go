package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Comment 레포지토리 인터페이스
type CommentRepository interface {
	FindAll() (res []entity.Comment, err error)
	Create(req *entity.Comment) (err error)
	Update(req *entity.Comment) (err error)
	Delete(id int) (res *entity.Comment, err error)
}

type gormCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &gormCommentRepository{db}
}

// 혈서목록조회
func (g *gormCommentRepository) FindAll() (res []entity.Comment, err error) {
	// 1. 쿼리 작성
	// select (id, dex_id, user_id, content) from comment

	// 2. gorm 적용
	tx := g.db
	tx.Model(&entity.Comment{}).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

// 혈서등록
func (g *gormCommentRepository) Create(userId uint64, content string) (err error) {
	// 1. 쿼리 작성
	// insert into comment (id, dex_id, user_id, content)
	// values (1, 1, 1, "내용")

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Create(&entity.Comment{
		UserId:  int(userId),
		Content: content,
	})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return tx.Error
}

// 혈서수정
func (g *gormCommentRepository) Update(id uint64, content string) (err error) {
	// 1. 쿼리 작성
	// update "comment"
	//    set content = '수정내용'
	//  where content is not null;

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Where("id = ?", id).Updates(content)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()

	return err
}

// 혈서삭제
func (g *gormCommentRepository) Delete(id int) (res *entity.Comment, err error) {
	// 1. 쿼리 작성
	// delete
	//   from "comment"
	//  where id = 4;

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Where("id = ?", id).Find(&res)
	tx.Delete(&res)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return res, nil
}
