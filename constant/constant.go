package constant

import (
	"main/constant/path/core"
	"sync"
)

var (
	instance *core.InternalApi
	once     sync.Once
)

// 라우터 path 함수
func GetCommentPath() *core.InternalApi {
	// Do함수는 인스턴스가 있으면 실행X
	once.Do(func() {
		instance = &core.InternalApi{
			Comment: core.CommentPath{
				GetAllComment: "/:id/comments",
				CreateComment: "/:id/comments",
				UpdateComment: "/:id/comments",
				DeleteComment: "/:id/comments",
			},
		}
	})

	return instance
}

func GetDexPath() *core.InternalApi {
	// Do함수는 인스턴스가 있으면 실행X
	once.Do(func() {
		instance = &core.InternalApi{
			Dex: core.DexPath{
				GetDex:    "/", // gRPC?
				CreateDex: "/",
			},
		}
	})

	return instance
}

// Do의 역할 instance 안에 내용이 없을 때 실행된다. (싱글톤패턴)
