package routes

import (
	"auth-system/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/Login", handlers.Login)
}