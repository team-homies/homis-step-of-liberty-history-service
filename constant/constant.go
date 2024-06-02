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
	once.Do(func() {
		instance = &core.InternalApi{
			Dex: core.DexPath{
				FindDexEvent:   "/history/:id",
				CreateDexEvent: "/history",
				GetTags:        "/tags",
			},
		}
	})

	return instance
}
