package dex

import (
	"main/app/api/dex/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

func SetDexApis(route fiber.Router) {
	h := handler.NewDexHandler()
	// 도감 필터 조회
	route.Get(constant.GetPath().Dex.GetTags, h.GetTags)
}
