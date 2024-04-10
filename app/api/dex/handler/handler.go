package handler

import (
	"main/app/api/dex/resource"
	"main/app/api/dex/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

// handler 인터페이스
type handler interface {
	GetDex(c *fiber.Ctx) error
	GetDexs(c *fiber.Ctx) error
	CreateDex(c *fiber.Ctx) error
	UpdateDex(c *fiber.Ctx) error
	DeleteDex(c *fiber.Ctx) error
}

// 핸들러 자료형?
type dexHandler struct {
	// 필요한 리소스 선언
	service service.DexService
}

func NewDexHandler() handler {
	return &dexHandler{
		// dex 서비스 생성
		service: service.NewDexService(),
	}
}

// comment.go에서 라우터에 들어가는 함수
// (h *commentHandler) 리시버함수를 만드는 방법
// GetComment 리시버함수
func (h *dexHandler) GetDex(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}     // 파이버객체 생성
	req := new(resource.GetDexRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.QueryParser(req)             // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.GetDex(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}

func (h *dexHandler) GetDexs(c *fiber.Ctx) error {
	return nil
}

func (h *dexHandler) CreateDex(c *fiber.Ctx) error {
	return nil
}

func (h *dexHandler) UpdateDex(c *fiber.Ctx) error {
	return nil
}

func (h *dexHandler) DeleteDex(c *fiber.Ctx) error {
	return nil
}
