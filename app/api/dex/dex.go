package dex

import (
	"main/app/api/dex/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

// router에 들어가는 함수
func SetDexApis(route fiber.Router) {
	h := handler.NewDexHandler()

	// 사건 내용 조회
	route.Get(constant.DexPath().DexEvent.FindDexEvent, h.FindDexEvent)
	// 사용자 사건 수집 등록
	route.Post(constant.DexPath().DexEvent.CreateDexEvent, middleware.JWTMiddleware, h.CreateDexEvent)
	// 도감 필터 조회
	route.Get(constant.GetPath().Dex.GetTags, h.GetTags)
}
