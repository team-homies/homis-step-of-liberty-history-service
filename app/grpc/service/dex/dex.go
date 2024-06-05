package dex

import (
	"context"
	"main/app/grpc/proto/dex"
	"main/database/repository"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	dex.DexEventServiceServer
}

func RegisterDosageService(grpcServer *grpc.Server) {
	dex.RegisterDexEventServiceServer(grpcServer, &server{})
}

// 사건 내용 조회
func (s *server) FindDexEvent(ctx context.Context, in *dex.DexEventRequest) (res *dex.DexEventResponse, err error) {
	dexReposiroty := repository.NewRepository()
	res = new(dex.DexEventResponse)

	// 1. 만들어진 레포지토리 두개를 사용해서 각각 데이터를 가져온다
	dexEvent, err := dexReposiroty.FindDexEventByEventId(int(in.EventId))
	if err != nil {
		return nil, err
	}
	dexDetail, err := dexReposiroty.FindDexDetailByEventId(int(in.EventId))
	if err != nil {
		return nil, err
	}

	// 2. 가져온 데이터를 하나의 객체(res)에 합친다

	res = &dex.DexEventResponse{
		Name:  dexEvent.Name,
		Level: dexEvent.Level,
		EventDetail: &dex.Detail{
			Define:     dexDetail.Define,
			Outline:    dexDetail.Outline,
			Place:      dexDetail.Place,
			Background: dexDetail.Background,
			ImageUrl:   dexDetail.ImageUrl,
		},
	}

	// 3. 리턴한다
	return
}

func (s *server) GetRate(ctx context.Context, in *dex.RateRequest) (*dex.RateResponse, error) {
	// 1. CountEvents 가져오고
	e, err := repository.NewRepository().CountEvents()
	if err != nil {
		return nil, err
	}
	// 2. CountUserEvents 가져와서
	u, err := repository.NewRepository().CountUserEvents(in.UserId)
	if err != nil {
		return nil, err
	}
	// u가 0일 때 그.. rate 계산이 안됨
	if u == 0 {
		return &dex.RateResponse{Rate: "0"}, err
	}
	// 3. CountUserEvents / CountEvents * 100
	var rate float64 = float64(u) / float64(e) * 100

	return &dex.RateResponse{
		Rate: strconv.FormatFloat(rate, 'f', -1, 64),
	}, err
}
