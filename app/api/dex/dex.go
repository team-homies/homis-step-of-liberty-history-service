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
	route.Get(constant.DexPath().Dex.FindDexEvent, h.FindDexEvent)
	// 도감 필터 조회
	route.Get(constant.DexPath().Dex.GetTags, h.GetTags)
	// 사용자 사건 수집 등록
	route.Post(constant.DexPath().Dex.CreateDexEvent, middleware.AuthVerificationMiddleware, h.CreateDexEvent)
	route.Get(constant.GetPath().Dex.GetTags, h.GetTags)
	// 명언 조회
	route.Get(constant.GetPath().Dex.GetQuote, h.GetQuote)
}
