package handler

import (
	"main/app/api/dex/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetTags(c *fiber.Ctx) error
	GetQuote(c *fiber.Ctx) error
}

type dexHandler struct {
	service service.DexService
}

func NewDexHandler() handler {
	return &dexHandler{
		service: service.NewDexService(),
	}
}
// 도감 필터 조회
func (h *dexHandler) GetTags(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	res, err := h.service.GetTags()
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	return ctx.HttpOK(res)
}
// 명언 조회
func (h *dexHandler) GetQuote(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	res, err := h.service.GetQuote()
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
