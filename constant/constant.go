package constant

import (
	"main/constant/path/core"
	"sync"
)

var (
	instance *core.InternalApi
	once     sync.Once
)

func GetPath() *core.InternalApi {
	once.Do(func() {
		instance = &core.InternalApi{
			Dex: core.DexPath{
				GetTags:  "/tags",
				GetQuote: "/quotes",
			},
		}
	})

	return instance
}
