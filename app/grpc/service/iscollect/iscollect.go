package iscollect

import (
	"context"
	"main/app/grpc/proto/dex"
	"main/app/grpc/proto/iscollect"
	"main/database/repository"

	"google.golang.org/grpc"
)

type server struct {
	iscollect.IsCollectServiceServer
	dex.DexEventServiceServer
}

// 사건 보유 여부 grpc
func FindIsCollectEventByUserId(ctx context.Context, in *iscollect.IsCollectRequest) (*iscollect.IsCollectResponse, error) {
	// 1. userid와 eventid일치 카운트 함수 호출
	decCount, err := repository.NewRepository().FindUserDexById(int(in.EventId), int(in.UserId))
	if err != nil {
		return nil, err
	}
	// 2. 보유 여부 조건식
	if decCount != 0 {
		// 보유한 사건이 0이 아니라면
		return &iscollect.IsCollectResponse{
			Division: "ture",
		}, err
	} else {
		// 보유한 사건이 0이라면
		return &iscollect.IsCollectResponse{
			Division: "false",
		}, err
	}

}

func RegisterDosageService(grpcServer *grpc.Server) {
	iscollect.RegisterIsCollectServiceServer(grpcServer, &server{})
}
