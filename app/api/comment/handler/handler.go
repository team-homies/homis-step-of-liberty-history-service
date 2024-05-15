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
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// 혈서수정
func (h *commentHandler) UpdateComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	id := c.Params("id")
	num, _ := strconv.Atoi(id)
	req := new(resource.UpdateCommentRequest)
	ctx.C.ParamsParser(req)
	err := h.service.UpdateComment(num, req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// 혈서삭제
func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.DeleteCommentRequest)
	id := c.Params("id")
	num, _ := strconv.Atoi(id)
	ctx.C.ParamsParser(req)
	res, err := h.service.DeleteComment(num)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
