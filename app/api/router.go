package api

import (
	"main/app/api/comment"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	commentGroup := app.Group("/history")
	comment.SetCommentApis(commentGroup)

}
