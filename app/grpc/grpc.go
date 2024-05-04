package grpc

import (
	"context"
	dex "main/app/grpc/proto/dex"
	dosage "main/app/grpc/proto/dosage"
	patient "main/app/grpc/proto/patient"

	"google.golang.org/grpc"
)

// gRPC함수를 가져다 쓴다는게
// grpc.go에서 pb.go의 함수를 가져와 쓴다는 말
type GrpcService interface {
	GetPatientById(ctx context.Context, in *patient.PatientRequest) (*patient.PatientResponse, error)
	GetDosageById(ctx context.Context, in *dosage.DosageRequest) (*dosage.DosageResponse, error)
	GetDexById(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error)
}

type grpcServer struct {
	patientService patient.UnimplementedPatientServiceServer
	dosageService  dosage.UnimplementedDosageServiceServer
	dex            dex.UnimplementedDexEventServiceServer
}

func NewGrpcServer() GrpcService {
	return &grpcServer{
		patient.UnimplementedPatientServiceServer{},
		dosage.UnimplementedDosageServiceServer{},
		dex.UnimplementedDexEventServiceServer{},
	}
}

func ListenGrpcServer() *grpc.Server {
	// 서비스 등록
	s := grpc.NewServer()
	patient.RegisterPatientServiceServer(s, patient.UnimplementedPatientServiceServer{})
	dosage.RegisterDosageServiceServer(s, dosage.UnimplementedDosageServiceServer{})
	dex.RegisterDexEventServiceServer(s, dex.UnimplementedDexEventServiceServer{})
	return s
}

func (s *grpcServer) GetPatientById(ctx context.Context, in *patient.PatientRequest) (*patient.PatientResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &patient.PatientResponse{
		PatientNo: in.GetPatientNo(),
		Name:      "안녕디지몬",
		Age:       24,
		BirthDate: "2000-09-07",
		Gender:    "unknown",
		Alived:    false,
	}, nil
}

func (s *grpcServer) GetDosageById(ctx context.Context, in *dosage.DosageRequest) (*dosage.DosageResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &dosage.DosageResponse{
		DosageNo: in.GetDosageNo(),
		DrugInfo: &dosage.Drug{
			Name:        "SA225P2",
			Description: "만병통치약",
			Usage:       "1일 1024회 복용",
			SideEffect:  "72시간 뒤 사망 할 수 있음",
		},
	}, nil
}

func (s *grpcServer) GetDexById(ctx context.Context, in *dex.DexEventRequest) (*dex.DexEventResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &dex.DexEventResponse{
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
