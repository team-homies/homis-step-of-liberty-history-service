package api

import (
	"main/app/api/dex"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	dexGroup := app.Group("")
	dex.SetDexApis(dexGroup)
}
