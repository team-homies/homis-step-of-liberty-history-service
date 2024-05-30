package repository

import (
	"main/database"
	"main/database/repository/comment"
	"main/database/repository/dex"
)

type Repository interface {
	dex.DexRepository
	comment.CommentRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		dex.NewDexRepository(db),
		comment.NewCommentRepository(db),
	}
}

type repository struct {
	dex.DexRepository
	comment.CommentRepository
}
