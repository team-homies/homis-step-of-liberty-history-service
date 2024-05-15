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

func DexPath() *core.InternalApi {
	// Do함수는 인스턴스가 있으면 실행X
	once.Do(func() {
		instance = &core.InternalApi{
			DexEvent: core.DexEventPath{
				FindDexEvent:   "/history",
				CreateDexEvent: "/history",
			},
		}
	})

	return instance
}

// Do의 역할 instance 안에 내용이 없을 때 실행된다. (싱글톤패턴)
