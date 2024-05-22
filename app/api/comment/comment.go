package comment

import (
	"main/app/api/comment/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

// router에 들어가는 함수
func SetCommentApis(route fiber.Router) {
	// /history
	h := handler.NewCommentHandler()

	// 혈서 list 목록
	route.Get(constant.CommentPath().Comment.FindAllComment, h.FindAllComment)
	// 혈서 등록
	route.Post(constant.CommentPath().Comment.CreateComment, middleware.JWTMiddleware, h.CreateComment)
	// 혈서 수정
	route.Put(constant.CommentPath().Comment.UpdateComment, middleware.JWTMiddleware, h.UpdateComment)
	// 혈서 제거
	route.Delete(constant.CommentPath().Comment.DeleteComment, middleware.JWTMiddleware, h.DeleteComment)
}
