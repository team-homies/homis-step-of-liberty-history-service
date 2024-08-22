package common

import (
	"context"
	"main/app/api/dex/resource"
	"main/app/grpc/proto/userlist"
	"main/config"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// 유저 리스트 grpc
func GetUserListGrpc(userId uint) (res *resource.GetRatesResponse, err error) {

	// 0. grpc 연동
	var address string
	if viper.GetString(config.GRPC_AUTH_HOST) == "localhost" {
		address = viper.GetString(config.GRPC_AUTH_HOST)
	} else {
		address = viper.GetString(config.GRPC_AUTH_HOST) + viper.GetString(config.GRPC_AUTH_HOST)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	userClient := userlist.NewUserListServiceClient(conn)

	// 1. userId를 이용해서 user의 수집률을 구하고 변수에 담는다
	userInfo, err := userClient.GetUserList(context.Background(), &userlist.UserListRequest{
		UserId: uint64(userId),
	})
	if err != nil {
		return
	}

	res = &resource.GetRatesResponse{
		Nickname: userInfo.Nickname,
		Rate:     userInfo.Rate,
	}

	return
}
