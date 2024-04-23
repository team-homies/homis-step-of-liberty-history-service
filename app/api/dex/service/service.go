package service

import "main/app/api/dex/resource"

// 서비스 인터페이스
type DexEventService interface {
	GetDexEvent() (res *resource.GetDexEventResponse, err error)
	CreateDexEvent(id uint, req *resource.CreateDexEventRequest) (err error)
}

// comment 서비스 함수
func NewDexEventService() DexEventService {
	return &dexEventService{
		// 위의 서비스 인터페이스 Comment서비스
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
func (s *dexEventService) CreateDexEvent(id uint, req *resource.CreateDexEventRequest) (err error) {
	for _, newDex := range req {
		dex := resource.CreateDexEventRequest{
			Id:      newDex.Id,
			Content: newDex.Content,
		}
		err = dex.DexEventRepository.Create(req)
		if err != nil {
			return err
		}
	}

	return nil
}
