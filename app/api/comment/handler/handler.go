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
	FindAllComment(c *fiber.Ctx) (err error)
	CreateComment(c *fiber.Ctx) (err error)
	UpdateComment(c *fiber.Ctx) (err error)
	DeleteComment(c *fiber.Ctx) (err error)
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

// [혈서 목록 조회] query : Id
func (h *commentHandler) FindAllComment(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}

	// 1. id값 받아오기
	req := new(resource.FindAllCommentRequest)
	idByParam := ctx.C.Params("id")
	req.Id, err = strconv.Atoi(idByParam)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 2. 서비스 함수 실행
	res, err := h.service.FindAllComment(req.Id)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

// [혈서 등록] query : Id, body : UserId, Content
func (h *commentHandler) CreateComment(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}

	req := new(resource.CreateCommentRequest)
	err = ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 1. Id값 받아오기
	idByParam := ctx.C.Params("id")
	req.Id, err = strconv.Atoi(idByParam)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 2. userId값 받아오기
	req.UserId = ctx.GetLocalsInt("user_id")

	// 3. 서비스 함수 실행
	err = h.service.CreateComment(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// [혈서 수정] query : Id, body : Id, Content
func (h *commentHandler) UpdateComment(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}

	req := new(resource.UpdateCommentRequest)
	err = ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 1. Id값 받아오기
	idByParam := ctx.C.Params("id")
	req.Id, err = strconv.Atoi(idByParam)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 2. userId값 받아오기
	req.UserId = ctx.GetLocalsInt("user_id")

	// 3. 서비스 함수 실행
	err = h.service.UpdateComment(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// [혈서 삭제] query : Id, body : Id
func (h *commentHandler) DeleteComment(c *fiber.Ctx) (err error) {
	ctx := fiberkit.FiberKit{C: c}

	req := new(resource.DeleteCommentRequest)
	ctx.C.BodyParser(req)

	// 1. Id값 받아오기
	idByParam := ctx.C.Params("id")
	req.Id, err = strconv.Atoi(idByParam)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 2. Id값 받아오기
	req.UserId = ctx.GetLocalsInt("user_id")

	// 3. 서비스 함수 실행
	err = h.service.DeleteComment(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}
