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

func (g *gormCommentRepository) FindAll() (res []entity.Comment, err error) {
	// 1. 쿼리 작성
	// select (id, dex_id, user_id, content) from comment

	// 2. gorm 적용
	tx := g.db
	tx.Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormCommentRepository) Create(req *entity.Comment) (err error) {
	// 1. 쿼리 작성
	// insert into comment (id, dex_id, user_id, content) values (1, 1, 1, "내용")

	// 2. gorm 적용
	tx := g.db.Begin()
	tx.Create(&req)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return tx.Error
}

func (g *gormCommentRepository) Update(req *entity.Comment) (err error) {
	// panic("")
	tx := g.db.Begin()
	tx.Model(&entity.Comment{}).Where("id = ?", req.ID).Updates(req)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()

	return err
}

func (g *gormCommentRepository) Delete(id int) (res *entity.Comment, err error) {
	// panic("")
	tx := g.db.Begin()
	tx.Where("id = ?", id).Find(&res)
	tx.Delete(&res)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}
	tx.Commit()
	return res, nil
}
