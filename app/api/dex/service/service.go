package service

import "main/app/api/dex/resource"

// 서비스 인터페이스
type DexService interface {
	GetDex() (res *resource.GetDexResponse, err error)
	CreateDex(id uint, req *resource.CreateDexEventRequest) (err error)
}

// comment 서비스 함수
func NewDexService() DexService {
	return &dexService{
		// 위의 서비스 인터페이스 Comment서비스
		DexService: &dexService{},
	}
}

type dexService struct {
	DexService
}

func (s *dexService) GetDex() (res *resource.GetDexResponse, err error) {
	res = new(resource.GetDexResponse)
	res = &resource.GetDexResponse{
		Dex: resource.DexEventResource{
			Id:       1,
			UserId:   1,
			UserName: "논현동",
			Content:  "time to go to bed",
		},
	}

	return
}
func (s *dexService) CreateDex(id uint, req *resource.CreateDexEventRequest) (err error) {
	for _, newDex := range req {
		dex := resource.CreateDexEventRequest{
			Id:      newDex.Id,
			Content: newDex.Content,
		}
		err = dex.DexRepository.Create(req)
		if err != nil {
			return err
		}
	}

	return nil
}
