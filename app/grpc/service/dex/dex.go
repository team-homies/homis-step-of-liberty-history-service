package dex

import (
	"context"
	"main/app/grpc/proto/dex"
	"main/database/repository"

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
	dexEvent, err := dexReposiroty.FindDexEventById(int(in.EventId))
	if err != nil {
		return nil, err
	}
	dexDetail, err := dexReposiroty.FindDexDetailById(int(in.EventId))
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
