package handler

import (
	"main/app/api/dex/resource"
	"main/app/api/dex/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetDexEvent(c *fiber.Ctx) error
	CreateDexEvent(c *fiber.Ctx) error
}

type dexEventHandler struct {
	service service.DexEventService
}

func NewDexEventHandler() handler {
	return &dexEventHandler{
		service: service.NewDexEventService(),
	}
}

func (h *dexEventHandler) GetDexEvent(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.GetDexEventRequest)
	ctx.C.QueryParser(req)
	res, err := h.service.GetDexEvent(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

func (h *dexEventHandler) CreateDexEvent(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.CreateDexEventRequest)
	ctx.C.ParamsParser(req)
	err := h.service.CreateDexEvent(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}
