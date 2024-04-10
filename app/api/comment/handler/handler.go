package handler

import (
	"main/app/api/comment/resource"
	"main/app/api/comment/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

// handler 인터페이스
type handler interface {
	GetComment(c *fiber.Ctx) error
	GetComments(c *fiber.Ctx) error
	CreateComment(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
}

// 핸들러 자료형?
type commentHandler struct {
	// 필요한 리소스 선언
	service service.CommentService
}

func NewCommentHandler() handler {
	return &commentHandler{
		// comment 서비스 생성
		service: service.NewCommentService(),
	}
}

// comment.go에서 라우터에 들어가는 함수
// (h *commentHandler) 리시버함수를 만드는 방법
// GetComment 리시버함수
func (h *commentHandler) GetComment(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}         // 파이버객체 생성
	req := new(resource.GetCommentRequest) // 사용자에게서 받은 요청값을 req에 받는다
	ctx.C.QueryParser(req)                 // 쿼리 | 제이슨은 바디파서 | 패스는 파람파서
	res, err := h.service.GetComment(req)

	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound) // 파이버키트에서 실패메세지 가져옴
	}
	return ctx.HttpOK(res) // 파이버키트에서 성공메세지 가져옴
}

func (h *commentHandler) GetComments(c *fiber.Ctx) error {
	return nil
}

func (h *commentHandler) CreateComment(c *fiber.Ctx) error {
	return nil
}

func (h *commentHandler) UpdateComment(c *fiber.Ctx) error {
	return nil
}

func (h *commentHandler) DeleteComment(c *fiber.Ctx) error {
	return nil
}
