package handler

import (
	"main/app/api/dex/service"
	"main/common/fiberkit"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
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
func (h *dexEventHandler) CreateDexEvent(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}
	// 1. id값 받아오기
	userId := c.Params("id")
	dexId := c.Params("id")
	userNum, _ := strconv.Atoi(userId)
	dexNum, _ := strconv.Atoi(dexId)
	ctx.C.ParamsParser(userId)
	ctx.C.ParamsParser(dexId)
	// 2. 서비스 함수 실행
	err = h.service.CreateUserDex(int64(userNum), int64(dexNum))

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}
