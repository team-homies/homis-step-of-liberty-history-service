package api

import (
	dex "main/app/api/dex"

	"github.com/gofiber/fiber/v2"
)

// 라우터 comment.SetApis
func InitRoutes(app *fiber.App) {
	dexApi := app.Group("")
	dex.SetDexApis(dexApi)
}
