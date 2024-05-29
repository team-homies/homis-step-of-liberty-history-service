package constant

import (
	"main/constant/path/core"
	"sync"
)

var (
	instance *core.InternalApi
	once     sync.Once
)

func DexPath() *core.InternalApi {
	once.Do(func() {
		instance = &core.InternalApi{
			DexEvent: core.DexEventPath{
				FindAllComment: "",
				CreateComment:  "",
				UpdateComment:  "",
				DeleteComment:  "",
				FindDexEvent:   "/history/:id",
				CreateDexEvent: "/history",
			},
		}
	})
	return instance
}
