package router

import (
	"github.com/VladRomanciuc/CSProject/auth/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Routes for fiber
func Routes(app *fiber.App) {
	api := app.Group("/auth", logger.New())
	api.Get("/facebook", handler.FacebookAuth)
	api.Get("/facebook/callback", handler.FacebookCallback)
	api.Get("/google", handler.GoogleAuth)
	api.Get("/google/callback", handler.GoogleCallback)
}