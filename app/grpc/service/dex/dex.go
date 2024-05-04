package dex

import (
	"main/app/grpc/resource"
	"main/database/repository"
)

type DexEventService interface {
	FindDexEvent(id int64) (res *resource.DexEventResponse, err error)
}

func NewDexEventService() DexEventService {
	return &dexEventService{
		DexEventService: &dexEventService{},
	}
}

type dexEventService struct {
	DexEventService
}

func (d *dexEventService) FindDexEvent(id int64) (res *resource.DexEventResponse, err error) {
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
	if err != nil {
		return nil, err
	}

	// 3. 리턴한다
	return res, nil
}
