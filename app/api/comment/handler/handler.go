package handler

import (
	"main/app/api/comment/resource"
	"main/app/api/comment/service"
	"main/common/fiberkit"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// 핸들러 인터페이스 선언 (메소드와 반환형 기재)
type handler interface {
	FindAllComment(c *fiber.Ctx) error
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

// 혈서목록조회
func (h *commentHandler) FindAllComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.GetAllCommentRequest)
	ctx.C.ParamsParser(req)
	res, err := h.service.FindAllComment()
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

// 혈서등록
func (h *commentHandler) CreateComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	id := c.Params("id")
	num, _ := strconv.Atoi(id)
	req := new(resource.CreateCommentRequest)
	ctx.C.ParamsParser(req)
	err := h.service.CreateComment(num, req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(err) // 파이버키트에서 성공메세지 가져옴
}

// 혈서수정
func (h *commentHandler) UpdateComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c} // 파이버객체 생성
	id := c.Params("id")
	num, _ := strconv.Atoi(id)                // 문자열id를 10진수 int로 변환
	req := new(resource.UpdateCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.ParamsParser(req)                   // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	err := h.service.UpdateComment(num, req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(err) // 파이버키트에서 성공메세지 가져옴
}

// 혈서삭제
func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}            // 파이버객체 생성
	req := new(resource.DeleteCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	id := c.Params("id")
	num, _ := strconv.Atoi(id) // 문자열id를 10진수 int로 변환
	ctx.C.ParamsParser(req)    // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.DeleteComment(num)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}
