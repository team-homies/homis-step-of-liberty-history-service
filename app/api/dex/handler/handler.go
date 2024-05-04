package handler

import (
	"main/app/api/dex/service"
)

type handler interface {
	// GetDexEvent(c *fiber.Ctx) error
	// CreateDexEvent(c *fiber.Ctx) error
}

type dexEventHandler struct {
	service service.DexEventService
}

func NewDexEventHandler() handler {
	return &dexEventHandler{
		service: service.NewDexEventService(),
	}
}

// func (h *dexEventHandler) CreateDexEvent(c *fiber.Ctx) error {
// 	ctx := fiberkit.FiberKit{C: c}
// 	req := new(resource.CreateDexEventRequest)
// 	ctx.C.ParamsParser(req)
// 	err := h.service.CreateDexEvent(req)

// 	if err != nil {
// 		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
// 	}
// 	return ctx.HttpOK(err)
// }
