package dex

import "main/app/api/dex/resource"

type DexEventService interface {
	GetDexEvent() (res *resource.GetDexEventResponse, err error)
}

func NewDexEventService() DexEventService {
	return &dexEventService{
		DexEventService: &dexEventService{},
	}
}

type dexEventService struct {
	DexEventService
}

func (s *dexEventService) GetDexEvent() (res *resource.GetDexEventResponse, err error) {
	res = new(resource.GetDexEventResponse)
	res = &resource.GetDexEventResponse{
		Dex: resource.DexEventResource{
			Id:       1,
			UserId:   1,
			UserName: "논현동",
			Content:  "time to go to bed",
		},
	}

	return
}
