package grpc

import (
	"context"
	dex "main/app/grpc/proto/dex"
	dexService "main/app/grpc/service/dex"

	"google.golang.org/grpc"
)

type GrpcService interface {
	FindDexEvent(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error)
}

type grpcServer struct {
	dex dex.UnimplementedDexEventServiceServer
}

func ListenGrpcServer() *grpc.Server {
	// 서비스 등록
	s := grpc.NewServer()

	dex.RegisterDexEventServiceServer(s, dex.UnimplementedDexEventServiceServer{})
	return s
}

func (s *grpcServer) FindDexEvent(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error) {
	// 1. id값 받아오기
	id := in.GetId()
	// 2. 서비스 함수 실행
	findDex, err := dexService.NewDexEventService().FindDexEvent(id)
	if err != nil {
		return nil, err
	}
	return &dex.DexEventResponse{
		Name:  findDex.Name,
		Level: findDex.Level,
		EventDetail: &dex.Detail{
			Define:     findDex.EventDetail.Define,
			Outline:    findDex.EventDetail.Outline,
			Place:      findDex.EventDetail.Place,
			Background: findDex.EventDetail.Background,
			ImageUrl:   findDex.EventDetail.ImageUrl,
		},
	}, nil
}
