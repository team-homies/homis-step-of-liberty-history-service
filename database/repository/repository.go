package repository

import (
	"main/database"
	"main/database/repository/dex"
)

type Repository interface {
	dex.DexRepository
}

type repository struct {
	dex.DexRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		dex.NewDexRepository(db),
	}
}
