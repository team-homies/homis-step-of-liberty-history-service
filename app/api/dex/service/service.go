package service

import "main/database/repository"

type DexEventService interface {
	CreateUserDex(dexId int64, userId int64) (err error)
}

func NewDexEventService() DexEventService {
	return &dexEventService{
		DexEventService: &dexEventService{},
	}
}

type dexEventService struct {
	DexEventService
}

func (s *dexEventService) CreateUserDex(dexId int64, userId int64) (err error) {
	dexReposiroty := repository.NewRepository()
	// 1. userId와 dexId가 일치하는 값 참거짓 구분
	countDex, err := dexReposiroty.FindUserDexById(dexId, userId)
	// 2. 만약 값이 0이 아니면 에러 반환
	if countDex != 0 {
		return err
		// 3. 만약 값이 0이면 Create 반환
	} else {
		err := dexReposiroty.CreateUserDexById(dexId, userId)
		if err != nil {
			return err
		}

	}
	return
}
