package middleware

import (
    "os"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    jwt "github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")

    if tokenString == "" {
        return fiber.NewError(fiber.StatusUnauthorized, "Token no proporcionado")
    }

    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    // Verificar el token JWT
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validar firma del token
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.ErrUnprocessableEntity
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Token inválido")
    }

    if !token.Valid {
        // Agregar validación de tiempo de expiración (10 minutos)
        claims := token.Claims.(jwt.MapClaims)
        exp := claims["exp"].(float64)
        if time.Now().Unix() > int64(exp) {
            return fiber.NewError(fiber.StatusUnauthorized, "Token expirado")
        }
    }

    // Si todo está bien, continuar con la ruta
    return c.Next()
}