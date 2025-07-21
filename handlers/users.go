package handlers

import (
	"api-rest-fiber/config"
	"api-rest-fiber/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    // Hashear la contraseña
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error al hashear la contraseña: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error interno del servidor",
        })
    }
    user.Password = string(hashedPassword)

    // Insertar el usuario en la base de datos
    collection := config.GetCollection("api_rest_fiber", "users")
    _, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        log.Printf("Error al insertar el usuario: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al registrar el usuario",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Usuario registrado exitosamente",
        "user":    user,
    })
}

func LoginUser(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	collection := config.GetCollection("api_rest_fiber", "users")

	err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Aquí generarías y devolverías un token JWT, por ahora solo un mensaje de éxito.
	return c.JSON(fiber.Map{"message": "Inicio de sesión exitoso"})
}

func GetUsers(c *fiber.Ctx) error {
    collection := config.GetCollection("api_rest_fiber", "users")

    // Consultar todos los usuarios
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Printf("Error al consultar usuarios: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al obtener los usuarios",
        })
    }
    defer cursor.Close(context.TODO())

    var users []models.User
    if err := cursor.All(context.TODO(), &users); err != nil {
        log.Printf("Error al decodificar usuarios: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al obtener los usuarios",
        })
    }

    return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
    id := c.Params("id") // Obtener el ID del usuario desde la URL

    collection := config.GetCollection("api_rest_fiber", "users")

    // Consultar el usuario por ID
    var user models.User
    err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
    if err != nil {
        log.Printf("Error al consultar el usuario: %v", err)
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Usuario no encontrado",
        })
    }

    return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id") // Obtener el ID del usuario desde la URL

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    collection := config.GetCollection("api_rest_fiber", "users")

    // Actualizar el usuario
    _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updateData})
    if err != nil {
        log.Printf("Error al actualizar el usuario: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al actualizar el usuario",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Usuario actualizado exitosamente",
    })
}

func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id") // Obtener el ID del usuario desde la URL

    collection := config.GetCollection("api_rest_fiber", "users")

    // Eliminar el usuario
    _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
    if err != nil {
        log.Printf("Error al eliminar el usuario: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error al eliminar el usuario",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Usuario eliminado exitosamente",
    })
}