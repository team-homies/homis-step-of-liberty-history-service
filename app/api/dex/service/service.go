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

func (s *dexService) GetComment(req *resource.GetDexEventRequest) (res *resource.GetDexResponse, err error) {
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
func (s *dexService) CreateComment(id uint, req *resource.CreateDexEventRequest) (err error) {
	for _, newComment := range req {
		comment := resource.CreateDexEventRequest{
			Id:      newComment.Id,
			Content: newComment.Content,
		}
		err = comment.CommentRepository.Create(req)
		if err != nil {
			return err
		}
	}

	return nil
}
