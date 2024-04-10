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

	// 혈서 목록
	//  /comments
	route.Get(constant.GetCommentPath().Comment.GetComment, h.GetComment)
	// 혈서 목록s
	route.Get(constant.GetCommentPath().Comment.GetComments, h.GetComments)
	// 혈서 등록
	route.Post(constant.GetCommentPath().Comment.CreateComment, h.CreateComment)
	// 혈서 수정
	route.Put(constant.GetCommentPath().Comment.UpdateComment, h.UpdateComment)
	// 혈서 제거
	route.Delete(constant.GetCommentPath().Comment.DeleteComment, h.DeleteComment)
}