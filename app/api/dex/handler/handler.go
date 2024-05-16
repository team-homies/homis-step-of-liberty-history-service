package handler

import (
	"main/app/api/dex/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetTags(c *fiber.Ctx) error
}

type dexHandler struct {
	service service.DexService
}

func NewDexHandler() handler {
	return &dexHandler{
		service: service.NewDexService(),
	}
}

func (h *dexHandler) GetTags(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	res, err := h.service.GetTags()
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	return ctx.HttpOK(res)
}
