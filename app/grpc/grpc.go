package grpc

import (
	"context"
	dex "main/app/grpc/proto/dex"
	dexService "main/app/grpc/service/dex"

	"google.golang.org/grpc"
)

// gRPC함수를 가져다 쓴다는게
// grpc.go에서 pb.go의 함수를 가져와 쓴다는 말
type GrpcService interface {
	FindDexEventById(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error)
}

type grpcServer struct {
	dex dex.UnimplementedDexEventServiceServer
}

func NewGrpcServer() GrpcService {
	return &grpcServer{
		dex.UnimplementedDexEventServiceServer{},
	}
}

func ListenGrpcServer() *grpc.Server {
	// 서비스 등록
	s := grpc.NewServer()

	dex.RegisterDexEventServiceServer(s, dex.UnimplementedDexEventServiceServer{})
	return s
}

func (s *grpcServer) FindDexEventById(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error) {
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
