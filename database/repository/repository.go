package repository

import (
	"main/database"
	"main/database/repository/comment"
)

type Repository interface {
	comment.CommentRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		comment.NewCommentRepository(db),
	}
}

type repository struct {
	comment.CommentRepository
}
