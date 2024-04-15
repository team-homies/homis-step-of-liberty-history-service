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
	// 도감 등록
	route.Post(constant.GetDexPath().Dex.CreateDex, h.CreateDex)
}
