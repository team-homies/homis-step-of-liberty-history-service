package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
)

type DexEventService interface {
	FindDexEvent(id int) (res *resource.FindEventResponse, err error)
	CreateUserDex(req *resource.CreateEventRequest) (err error)
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
func (s *dexEventService) CreateUserDex(req *resource.CreateEventRequest) (err error) {
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

// [사건 내용 조회] 사건 id로 조회 : 서비스
func (d *dexEventService) FindDexEvent(id int) (res *resource.FindEventResponse, err error) {
	dexReposiroty := repository.NewRepository()

	res = new(resource.FindEventResponse)

	// 1. 만들어진 레포지토리 두개를 사용해서 각각 데이터를 가져온다
	dexEvent, err := dexReposiroty.FindDexEventById(id)
	if err != nil {
		return nil, err
	}
	dexDetail, err := dexReposiroty.FindDexDetailById(id)
	if err != nil {
		return nil, err
	}

	// 2. 가져온 데이터를 하나의 객체(res)에 합친다

	res = &resource.FindEventResponse{
		Name:  dexEvent.Name,
		Level: dexEvent.Level,
		Detail: resource.FindDetailResponse{
			Define:     dexDetail.Define,
			Outline:    dexDetail.Outline,
			Place:      dexDetail.Place,
			Background: dexDetail.Background,
			ImageUrl:   dexDetail.ImageUrl,
		},
	}

	// 3. 리턴한다
	return res, nil
}
