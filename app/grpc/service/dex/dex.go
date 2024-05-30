package dex

import (
	"context"
	"main/app/grpc/proto/dex"

	"google.golang.org/grpc"
)

type server struct {
	dex.DexEventServiceServer
}

func RegisterDosageService(grpcServer *grpc.Server) {
	dex.RegisterDexEventServiceServer(grpcServer, &server{})
}

// proto에서 정의한 함수명과 동일하게 할 것
func (s *server) GetDexById(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &dex.DexEventResponse{
		EventId: in.GetId(),
		Name:    "3.1운동",
		Level:   "24",
		Created: "2024-05-01",
		Updated: "2024-05-03",
		IsUsed:  false,
		EventDetail: &dex.Detail{
			Define:     "1919년 3월 1일을 기해 일어난 거족적인 독립만세운동.",
			Outline:    "3·1운동은 1919년 3월 1일을 기해 일어난 거족적인 독립만세운동이다. 전국적인 범위에서 각계각층을 망라하여 전개된 3·1운동은 세계의 이목을 집중시켜 한국민에 대한 인식을 새롭게 하였고, 중국 상하이에서의 대한민국임시정부 수립으로 이어졌다. 이민족에 대한 끈질기고 강렬한 독립투쟁정신을 고취하였을 뿐 아니라, 일제의 무단통치방법을 이른바 문화통치로 바꾸게 하였다. 나아가 민족의식과 민족정신에 새로운 자각과 힘을 주어 교육의 진흥, 신문예운동·산업운동이 활성화하고 민족 자립의 기초를 다지게 하는 계기를 만들기도 하였다.",
			Place:      "한반도 전역",
			Background: "어쩌구",
			ImageUrl:   "image.url",
		},
	}, nil
}
