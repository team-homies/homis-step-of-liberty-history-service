package grpc

import (
	dex "main/app/grpc/service/dex"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	dex.RegisterDosageService(grpcServer)
}
