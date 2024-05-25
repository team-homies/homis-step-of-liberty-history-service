// grpc핸들러
// 2024.05.15 사건 내용 조회의 grpc 잠시 보류하기로 함
package grpc

import (
	"context"
	dex "main/app/grpc/proto/dex"

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

/*
// [사건 내용 조회] 사건 id로 조회 : 핸들러
func (s *grpcServer) FindDexEvent(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error) {
	// 1. id값 받아오기
	id := in.GetId()
	// 2. 서비스 함수 실행
	findDex, err := dexService.NewDexEventService().FindDexEvent(int(id))
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
*/
