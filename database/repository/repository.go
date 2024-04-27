package repository

import (
	"main/database"
)

type Repository interface {
	dex.DexRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		dex.NewDexRepository(db),
	}
}

type repository struct {
	dex.DexRepository
}
