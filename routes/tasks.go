package routes

import (
	"api-rest-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTasksRoutes(app *fiber.App) {
    api := app.Group("/api/v1/tasks")

    api.Post("/create", handlers.CreateTask)
    api.Delete("/delete/:id", handlers.DeleteTask)
    api.Put("/update/:id", handlers.UpdateTask)
}