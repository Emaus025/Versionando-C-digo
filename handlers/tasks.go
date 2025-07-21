package handlers

import (
	"api-rest-fiber/config"
	"api-rest-fiber/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTask crea una nueva tarea
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	collection := config.GetCollection("api_rest_fiber", "tasks")
	result, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear la tarea"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Tarea creada exitosamente", "taskID": result.InsertedID})
}

// UpdateTask actualiza una tarea existente
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de tarea inválido"})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	collection := config.GetCollection("api_rest_fiber", "tasks")
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": task})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al actualizar la tarea"})
	}

	return c.JSON(fiber.Map{"message": "Tarea actualizada exitosamente"})
}

// DeleteTask elimina una tarea
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID de tarea inválido"})
	}

	collection := config.GetCollection("api_rest_fiber", "tasks")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar la tarea"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarea no encontrada"})
	}

	return c.JSON(fiber.Map{"message": "Tarea eliminada exitosamente"})
}