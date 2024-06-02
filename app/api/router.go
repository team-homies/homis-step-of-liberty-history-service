package api

import (
	"main/app/api/comment"
	dex "main/app/api/dex"
	"main/app/api/dex"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	dexGroup := app.Group("")
	dex.SetDexApis(dexGroup)
	commentGroup := app.Group("/history")
	comment.SetCommentApis(commentGroup)
}
