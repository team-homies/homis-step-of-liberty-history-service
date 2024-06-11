package grpc

import (
	dex "main/app/grpc/service/dex"
	iscollect "main/app/grpc/service/iscollect"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	dex.RegisterService(grpcServer)
	iscollect.RegisterService(grpcServer)
}
