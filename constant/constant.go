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
				GetComment:    "/comment",
				GetComments:   "/comment/list",
				CreateComment: "/comment",
				UpdateComment: "/comment",
				DeleteComment: "/comment",
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
				GetDex:    "/dex",
				GetDexs:   "/dex/list",
				CreateDex: "/dex",
				UpdateDex: "/dex",
				DeleteDex: "/dex",
			},
		}
	})

	return instance
}

// Do의 역할 instance 안에 내용이 없을 때 실행된다. (싱글톤패턴)
