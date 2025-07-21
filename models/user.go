package models

import "time"

// User representa la estructura de un usuario
type User struct {
    ID             string    `json:"id,omitempty"`
    Name           string    `json:"name"`
    LastName       string    `json:"lastName"`
    Email          string    `json:"email"`
    Password       string    `json:"password"` // En producción, se debe almacenar hash de la contraseña
    BirthDate      time.Time `json:"birthDate"`
    SecretQuestion string    `json:"secretQuestion"`
    SecretAnswer   string    `json:"secretAnswer"`
}

