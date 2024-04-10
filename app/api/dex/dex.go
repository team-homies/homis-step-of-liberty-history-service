package comment

import (
	"main/app/api/dex/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

// router에 들어가는 함수
func SetDexApis(route fiber.Router) {
	h := handler.NewDexHandler()

	// 도감 목록
	route.Get(constant.GetDexPath().Dex.GetDex, h.GetDex)
	// 도감 목록s
	route.Get(constant.GetDexPath().Dex.GetDexs, h.GetDexs)
	// 도감 등록
	route.Post(constant.GetDexPath().Dex.CreateDex, h.CreateDex)
	// 도감 수정
	route.Put(constant.GetDexPath().Dex.UpdateDex, h.UpdateDex)
	// 도감 제거
	route.Delete(constant.GetDexPath().Dex.DeleteDex, h.DeleteDex)
}
