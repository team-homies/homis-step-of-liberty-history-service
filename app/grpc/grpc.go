package grpc

import (
	dex "main/app/grpc/service/dex"
	iscollect "main/app/grpc/service/iscollect"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	dex.RegisterDosageService(grpcServer)
	iscollect.RegisterDosageService(grpcServer)
}
