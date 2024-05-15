package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
)

type DexEventService interface {
	CreateUserDex(req *resource.CreateDexEventRequest) (err error)
}

func NewDexEventService() DexEventService {
	return &dexEventService{
		DexEventService: &dexEventService{},
	}
}

type dexEventService struct {
	DexEventService
}

// [사용자 사건 수집 등록] 사건 id로 등록 post문 : 서비스
func (s *dexEventService) CreateUserDex(req *resource.CreateDexEventRequest) (err error) {
	dexReposiroty := repository.NewRepository()
	// 1. userId와 dexId가 일치하는 값 참거짓 구분
	countDex, err := dexReposiroty.FindUserDexById(req.EventId, req.UserId)
	// 2. 만약 값이 0이 아니면 에러 반환
	if countDex != 0 {
		return err
		// 3. 만약 값이 0이면 Create 반환
	} else {
		err := dexReposiroty.CreateUserDexById(req.EventId, req.UserId)
		if err != nil {
			return err
		}

	}
	return
}
