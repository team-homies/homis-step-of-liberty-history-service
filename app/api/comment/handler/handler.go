package handler

import (
	"main/app/api/comment/resource"
	"main/app/api/comment/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

// 핸들러 인터페이스 선언 (메소드와 반환형 기재)
type handler interface {
	GetAllComment(c *fiber.Ctx) error
	CreateComment(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
}

// commentHandler형 선언
type commentHandler struct {
	// 필요한 리소스 선언
	service service.CommentService
}

// 함수를 handler형으로 선언
func NewCommentHandler() handler {
	// commentHandler의 주소값 반환
	return &commentHandler{
		service: service.NewCommentService(),
	}
}

// comment.go에서 라우터에 들어가는 함수
// (h *commentHandler) 리시버함수를 만드는 방법
// GetAllComment 리시버함수

func (h *commentHandler) GetAllComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}            // 파이버객체 생성
	req := new(resource.GetAllCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                   // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.GetAllComment()

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}

func (h *commentHandler) CreateComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}            // 파이버객체 생성
	req := new(resource.CreateCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                   // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	err := h.service.CreateComment(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(err) // 파이버키트에서 성공메세지 가져옴
}

func (h *commentHandler) UpdateComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}            // 파이버객체 생성
	req := new(resource.UpdateCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                   // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	err := h.service.UpdateComment(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(err) // 파이버키트에서 성공메세지 가져옴
}

func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}            // 파이버객체 생성
	req := new(resource.DeleteCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                   // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.DeleteComment(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}
