package comment

import (
	"main/database/entity"

	"gorm.io/gorm"
)

// Comment 레포지토리 인터페이스
type CommentRepository interface {
	GetAll() (res []entity.Comment, err error)
	Create(req *entity.Comment) (err error)
	Update(req *entity.Comment) (err error)
	Delete(id uint) (res *entity.Comment, err error)
}

type gormCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &gormCommentRepository{db}
}

func (g *gormCommentRepository) GetAll() (res []entity.Comment, err error) {
	// panic("error!!")
	tx := g.db
	tx.Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (g *gormCommentRepository) Create(req *entity.Comment) (err error) {
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

func (g *gormCommentRepository) Delete(id uint) (res *entity.Comment, err error) {
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
