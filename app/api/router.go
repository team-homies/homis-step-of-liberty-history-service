package api

import (
	"main/app/api/comment"

	"github.com/gofiber/fiber/v2"
)

// 라우터 comment.SetApis
func InitRoutes(app *fiber.App) {
	ptientGroup := app.Group("/api")
	comment.SetApis(ptientGroup)
	// ㄴ comment.SetApis(app.Group("/api"))
}
