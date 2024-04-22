package comment

import (
	"main/app/api/comment/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

// router에 들어가는 함수
func SetCommentApis(route fiber.Router) {
	// /history
	h := handler.NewCommentHandler()

	// 혈서 list 목록
	route.Get(constant.GetCommentPath().Comment.GetAllComment, h.GetAllComment)
	// 혈서 등록
	route.Post(constant.GetCommentPath().Comment.CreateComment, h.CreateComment)
	// 혈서 수정
	route.Put(constant.GetCommentPath().Comment.UpdateComment, h.UpdateComment)
	// 혈서 제거
	route.Delete(constant.GetCommentPath().Comment.DeleteComment, h.DeleteComment)
}
