package comment

import (
	"main/app/api/comment/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetCommentApis(route fiber.Router) {
	// /history
	h := handler.NewCommentHandler()

	// 혈서 목록 조회
	route.Get(constant.CommentPath().Comment.FindAllComment, h.FindAllComment)
	// 혈서 등록
	route.Post(constant.CommentPath().Comment.CreateComment, middleware.AuthVerificationMiddleware, h.CreateComment)
	// 혈서 수정
	route.Put(constant.CommentPath().Comment.UpdateComment, middleware.AuthVerificationMiddleware, h.UpdateComment)
	// 혈서 제거
	route.Delete(constant.CommentPath().Comment.DeleteComment, middleware.AuthVerificationMiddleware, h.DeleteComment)
}
