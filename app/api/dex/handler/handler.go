package handler

import (
	"main/app/api/dex/resource"
	"main/app/api/dex/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

// handler 인터페이스
type handler interface {
	GetDexEvent(c *fiber.Ctx) error
	CreateDexEvent(c *fiber.Ctx) error
}

// 핸들러 자료형?
type dexEventHandler struct {
	// 필요한 리소스 선언
	service service.DexEventService
}

func NewDexEventHandler() handler {
	return &dexEventHandler{
		// dex 서비스 생성
		service: service.NewDexEventService(),
	}
}

// comment.go에서 라우터에 들어가는 함수
// (h *commentHandler) 리시버함수를 만드는 방법
// GetComment 리시버함수
func (h *dexEventHandler) GetDexEvent(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}          // 파이버객체 생성
	req := new(resource.GetDexEventRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.QueryParser(req)                  // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.GetDexEvent()

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}

func (h *dexEventHandler) CreateDexEvent(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}             // 파이버객체 생성
	req := new(resource.CreateDexEventRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                    // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	err := h.service.CreateDexEvent(uint(req.Id), req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(err) // 파이버키트에서 성공메세지 가져옴
}