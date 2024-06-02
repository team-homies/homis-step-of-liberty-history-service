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
func CommentPath() *core.InternalApi {
	// Do함수는 인스턴스가 있으면 실행X
	once.Do(func() {
		instance = &core.InternalApi{
			Comment: core.CommentPath{
				FindAllComment: "/:id/comments",
				CreateComment:  "/:id/comments",
				UpdateComment:  "/:id/comments",
				DeleteComment:  "/:id/comments",
			},
		}
	})

	return instance
}

func DexPath() *core.InternalApi {
	// Do함수는 인스턴스가 있으면 실행X
	once.Do(func() {
		instance = &core.InternalApi{}
			Dex: core.DexPath{
				GetTags: "/tags",
				FindDexEvent:   "/history/:id",
				CreateDexEvent: "/history",
			},		
	})

	return instance
}
