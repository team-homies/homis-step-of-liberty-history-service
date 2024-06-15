package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
)

type DexService interface {
	GetTags() (res []resource.GetTagsResponse, err error)
	FindDexEvent(id int) (res *resource.FindEventResponse, err error)
	CreateUserDex(req *resource.CreateEventRequest) (err error)
}

func NewDexService() DexService {
	return &dexService{
		DexService: &dexService{},
	}
}

type dexService struct {
	DexService
}

// [사용자 사건 수집 등록] 사건 id로 등록 post문 : 서비스
func (s *dexService) CreateUserDex(req *resource.CreateEventRequest) (err error) {
	dexReposiroty := repository.NewRepository()
	// 1. userId와 dexId가 일치하는 값 참거짓 구분
	countDex, err := dexReposiroty.FindUserDexByEventId(req.EventId, req.UserId)
	// 2. 만약 값이 0이 아니면 에러 반환
	if countDex != 0 {
		return
		// 3. 만약 값이 0이면 Create 반환
	} else {
		err = dexReposiroty.CreateUserDexByEventId(req.EventId, req.UserId)
		if err != nil {
			return
		}

	}
	return
}

// [사건 내용 조회] 사건 id로 조회 : 서비스
func (d *dexService) FindDexEvent(id int) (res *resource.FindEventResponse, err error) {
	dexReposiroty := repository.NewRepository()

	res = new(resource.FindEventResponse)

	// 1. 만들어진 레포지토리 두개를 사용해서 각각 데이터를 가져온다
	dexEvent, err := dexReposiroty.FindDexEventByEventId(id)
	if err != nil {
		return
	}

	// 2. 가져온 데이터를 하나의 객체(res)에 합친다

	res = &resource.FindEventResponse{
		Name:  dexEvent.Name,
		Level: dexEvent.Level,
		Detail: resource.FindDetailResponse{
			Define:     dexEvent.Define,
			Outline:    dexEvent.Outline,
			Place:      dexEvent.Place,
			Background: dexEvent.Background,
			ImageUrl:   dexEvent.ImageUrl,
		},
	}

	return
}

func (d *dexService) GetTags() (res []resource.GetTagsResponse, err error) {
	res = []resource.GetTagsResponse{}
	tags, err := repository.NewRepository().GetTags()
	if err != nil {
		return
	}

	for _, tag := range tags {
		res = append(res, resource.GetTagsResponse{
			Id:   tag.ID,
			Name: tag.Name,
		})
	}
	return res, err
}
