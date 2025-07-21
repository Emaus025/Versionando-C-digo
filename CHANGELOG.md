# Changelog

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),
y este proyecto adhiere al [Versionado Semántico](https://semver.org/spec/v2.0.0.html).

## [Sin versionar] - 2024-12-19

### Agregado
- Configuración inicial del proyecto con Go Fiber
- Integración con MongoDB para persistencia de datos
- Sistema de autenticación JWT
- Middleware de autenticación para rutas protegidas
- Gestión de variables de entorno con godotenv
- Estructura modular del proyecto (handlers, models, routes, middleware)

#### Modelos
- Modelo `User` con campos: ID, Name, LastName, Email, Password, BirthDate, SecretQuestion, SecretAnswer
- Modelo `Task` con campos: ID, Title, Description, StartAt, DueDate, UserID

#### Endpoints de Usuarios
- `POST /api/v1/users/register` - Registro de nuevos usuarios
- `POST /api/v1/users/login` - Autenticación de usuarios

#### Endpoints de Tareas
- `POST /api/v1/tasks/create` - Crear nueva tarea (requiere autenticación)
- `PUT /api/v1/tasks/update/:id` - Actualizar tarea existente (requiere autenticación)
- `DELETE /api/v1/tasks/delete/:id` - Eliminar tarea (requiere autenticación)

#### Configuración
- Configuración de base de datos MongoDB
- Middleware de autenticación JWT
- Configuración de rutas y grupos de API
- Archivo .gitignore para excluir archivos sensibles

#### Dependencias
- `github.com/gofiber/fiber/v2 v2.52.0` - Framework web
- `github.com/golang-jwt/jwt/v4 v4.5.0` - Manejo de JWT
- `github.com/joho/godotenv v1.5.1` - Variables de entorno
- `go.mongodb.org/mongo-driver v1.13.1` - Driver de MongoDB
- `golang.org/x/crypto v0.33.0` - Funciones criptográficas

### Seguridad
- Implementación de encriptación de contraseñas
- Autenticación basada en JWT
- Middleware de autorización para rutas protegidas
- Gestión segura de variables de entorno

---

## Formato de Versiones Futuras

### [X.Y.Z] - YYYY-MM-DD

#### Agregado
- Nuevas características

#### Cambiado
- Cambios en funcionalidades existentes

#### Obsoleto
- Características que serán removidas en versiones futuras

#### Removido
- Características removidas

#### Corregido
- Corrección de bugs

#### Seguridad
- Mejoras de seguridad