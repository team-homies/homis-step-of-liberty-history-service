package handler

import (
	"main/app/api/dex/resource"
	"main/app/api/dex/service"
	"main/app/grpc/service/dex"
	"main/common/fiberkit"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	CreateDexEvent(c *fiber.Ctx) error
	FindDexEvent(c *fiber.Ctx) error
}

type dexEventHandler struct {
	service service.DexEventService
	dex     dex.DexEventService
}

func NewDexEventHandler() handler {
	return &dexEventHandler{
		service: service.NewDexEventService(),
	}
}

// [사용자 사건 수집 등록] 사건 id로 등록 post문 : 핸들러
func (h *dexEventHandler) CreateDexEvent(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}
	// 1. id값 받아오기
	req := new(resource.CreateDexEventRequest)
	ctx.C.BodyParser(req)

	req.UserId = ctx.GetLocalsInt("user_id")

	// 2. 서비스 함수 실행
	err = h.service.CreateUserDex(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// [사건 내용 조회] 사건 id로 조회 : 핸들러
func (h *dexEventHandler) FindDexEvent(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}
	// 1. id값 받아오기
	req := new(resource.FindDexEventRequest)
	findEventById := ctx.C.Params("id")
	req.EventId, _ = strconv.Atoi(findEventById)

	// 2. 서비스 함수 실행
	res, err := h.dex.FindDexEvent(req.EventId)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
