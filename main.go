package main

import (
    "github.com/joho/godotenv"
    "log"

    "github.com/gofiber/fiber/v2"
    "api-rest-fiber/config"
    "api-rest-fiber/routes"
    "api-rest-fiber/middleware"
)

func main() {
    // Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error al cargar el archivo .env")
    }

    app := fiber.New()

    // Conectar a MongoDB
    config.ConnectToDatabase()

    // Definir rutas
    routes.SetupUserRoutes(app)
    routes.SetupTasksRoutes(app) // Asegúrate de que esta línea esté presente

    // Agregar middleware global para rutas protegidas
    app.Use("/api", middleware.Auth)

    // Iniciar el servidor
    log.Fatal(app.Listen(":3000"))
}