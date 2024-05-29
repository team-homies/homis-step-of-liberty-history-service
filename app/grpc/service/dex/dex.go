// grpc 서비스
// 2024.05.15 사건 내용 조회의 grpc 잠시 보류하기로 함
package dex

import (
	"main/app/grpc/resource"
)

type DexEventService interface {
	FindDexEvent(id int) (res *resource.DexEventResponse, err error)
}

func NewDexEventService() DexEventService {
	return &dexEventService{
		DexEventService: &dexEventService{},
	}
}

type dexEventService struct {
	DexEventService
}

/*
// [사건 내용 조회] 사건 id로 조회 : 서비스
func (d *dexEventService) FindDexEvent(id int) (res *resource.DexEventResponse, err error) {
	dexReposiroty := repository.NewRepository()

	res = new(resource.DexEventResponse)

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
	res.Name = dexEvent.Name
	res.Level = dexEvent.Level
	res.EventDetail.Define = dexDetail.Define
	res.EventDetail.Outline = dexDetail.Outline
	res.EventDetail.Place = dexDetail.Place
	res.EventDetail.Background = dexDetail.Background
	res.EventDetail.ImageUrl = dexDetail.ImageUrl

	// 3. 리턴한다
	return res, nil
}
*/



