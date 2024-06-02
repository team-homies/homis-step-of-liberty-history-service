package api

import (
	"main/app/api/comment"
	dex "main/app/api/dex"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	dexApi := app.Group("")
	dex.SetDexApis(dexApi)
	commentGroup := app.Group("/history")
	comment.SetCommentApis(commentGroup)
}
