// grpc핸들러
// 2024.05.15 사건 내용 조회의 grpc 잠시 보류하기로 함
package grpc

import (
	dex "main/app/grpc/service/dex"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	dex.RegisterDosageService(grpcServer)
}
