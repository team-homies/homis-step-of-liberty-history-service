package service

import "main/app/api/dex/resource"

// 서비스 인터페이스
type DexService interface {
	GetDex(req *resource.GetDexRequest) (res *resource.GetDexResponse, err error)
	GetDexs() (resource.GetDexsResponse, error)
	CreateDex(req *resource.CreateDexRequest) error
	UpdateDex(req *resource.UpdateDexRequest) (res *resource.UpdateDexResponse, err error)
	DeleteDex(id int) error
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

func (s *dexService) GetComment(req *resource.GetDexRequest) (res *resource.GetDexResponse, err error) {
	res = new(resource.GetDexResponse)
	res = &resource.GetDexResponse{
		Dex: resource.DexResource{
			Id:       1,
			UserId:   1,
			UserName: "논현동",
			Content:  "time to go to bed",
		},
	}

	return
}
