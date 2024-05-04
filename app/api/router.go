package api

import (
	"main/app/api/comment"

	"github.com/gofiber/fiber/v2"
)

// 라우터 comment.SetApis
func InitRoutes(app *fiber.App) {
	commentApi := app.Group("/history")
	comment.SetCommentApis(commentApi)
	// dexApi := app.Group("/history")
	// dex.SetDexEventApis(dexApi)
}
