package api

import (
	"main/app/api/comment"
	dex "main/app/api/dex"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	commentGroup := app.Group("/history")
	comment.SetCommentApis(commentGroup)
	dexApi := app.Group("")
	dex.SetDexApis(dexApi)
}
