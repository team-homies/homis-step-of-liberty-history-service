package service

import (
	"main/app/api/dex/resource"
	"main/database/repository"
	"math/rand"
	"time"
)

type DexService interface {
	GetQuote() (res *resource.GetQuoteResponse, err error)
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

// 도감 필터 조회
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

// 명언 조회
var quotesOrigin map[int]*resource.GetQuoteResponse
var flag int
var quoteOfTheWeek *resource.GetQuoteResponse

// 명언 조회
func (d *dexService) GetQuote() (res *resource.GetQuoteResponse, err error) {
	// 1. 현재 요일 구하기 (서울 기준)
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return
	}
	// now := time.Now() // Go Playground 에서는 항상 시각은 2009-11-10 23:00:00 +0000 UTC 에서 시작한다.
	// t := now.In(loc) // 그래서 서울 기준 UTC 세팅(+0900)
	t := time.Now().In(loc)     //그래서 서울 기준 UTC 세팅(+0900)
	weekDay := int(t.Weekday()) //일요일 : 0, 월요일 : 1
	// 2. flag가 일요일 and weekDay가 월요일
	if flag == 0 && weekDay == 1 {
		quotes := make(map[int]*resource.GetQuoteResponse)
		// 3. 명언2(quotesOrigin)이 0일때
		if len(quotesOrigin) == 0 {
			// 3-1. 전체 명언을 map에 담기
			quoteEntities, err2 := repository.NewRepository().GetQuote()
			if err2 != nil {
				return nil, err2
			}
			// quotes := make([]resource.GetQuoteResponse, 0)
			// 3-2 인덱스와 명언 (key/value)
			for i, q := range quoteEntities {
				quotes[i] = &resource.GetQuoteResponse{
					Id:       q.ID,
					Content:  q.Content,
					ImageUrl: q.ImageUrl,
				}
			}
		} else { // 4. 명언2(quotesOrigin)이 0이 아닐 때 명언1(quotes)에 명언2(quotesOrigin)넣기
			quotes = quotesOrigin
		}
		// 5. quotes의 길이에서 랜덤으로 숫자 하나 rq에 담기
		// 별도 변수에 고른 명언을 담기
		randomQuote := rand.Intn(len(quotes))

		// 6. 고른 명언 반환
		res = quotes[randomQuote]

		// 7. rq를 지우고 새 map에 담기
		delete(quotes, randomQuote)

		// 8. 명언2(quotesOrgin) 에 명언1(quotes) 담기
		quotesOrigin = quotes

		// 9. 전역변수에 res담기
		quoteOfTheWeek = res
		if quoteOfTheWeek == nil {
			return
		}
	}
	// flag에 weekDay 담아놓기
	flag = weekDay

	return quoteOfTheWeek, err

}

// 월요일이 아니면 랜덤 로직이 안돌게 하지만 월요일에 최초로 호출된 경우의 명언을 보여줘야해요
