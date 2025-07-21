package routes

import (
	"api-rest-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
    api := app.Group("/api/v1/users")

    api.Post("/register", handlers.CreateUser)
    api.Post("/login", handlers.LoginUser)
}