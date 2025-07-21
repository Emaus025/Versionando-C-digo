# API REST con Go Fiber

Una API REST desarrollada con Go y el framework Fiber para la gestión de usuarios y tareas, con autenticación JWT y base de datos MongoDB.

## 🚀 Características

- **Autenticación JWT**: Sistema de autenticación seguro con tokens JWT
- **Gestión de Usuarios**: Registro y login de usuarios
- **Gestión de Tareas**: CRUD completo para tareas de usuarios
- **Base de Datos**: Integración con MongoDB
- **Middleware**: Autenticación y autorización
- **Estructura Modular**: Código organizado en handlers, models, routes y middleware

## 🛠️ Tecnologías Utilizadas

- **Go 1.21**
- **Fiber v2** - Framework web rápido y minimalista
- **MongoDB** - Base de datos NoSQL
- **JWT** - Autenticación con tokens
- **bcrypt** - Encriptación de contraseñas
- **godotenv** - Gestión de variables de entorno

## 📋 Prerrequisitos

- Go 1.21 o superior
- MongoDB (local o MongoDB Atlas)
- Git

## 🔧 Instalación

1. **Clonar el repositorio**
   ```bash
   git clone <url-del-repositorio>
   cd api-rest-fiber
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Configurar variables de entorno**
   
   Crear un archivo `.env` en la raíz del proyecto:
   ```env
   MONGODB_URI=mongodb://localhost:27017/api-rest-fiber
   JWT_SECRET=tu_clave_secreta_jwt
   PORT=3000
   ```

4. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

   La API estará disponible en `http://localhost:3000`

## 📚 Endpoints de la API

### Usuarios

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| POST | `/api/v1/users/register` | Registrar nuevo usuario | No |
| POST | `/api/v1/users/login` | Iniciar sesión | No |

### Tareas

| Método | Endpoint | Descripción | Autenticación |
|--------|----------|-------------|---------------|
| POST | `/api/v1/tasks/create` | Crear nueva tarea | Sí |
| PUT | `/api/v1/tasks/update/:id` | Actualizar tarea | Sí |
| DELETE | `/api/v1/tasks/delete/:id` | Eliminar tarea | Sí |

## 📝 Ejemplos de Uso

### Registrar Usuario
```bash
curl -X POST http://localhost:3000/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan",
    "lastName": "Pérez",
    "email": "juan@example.com",
    "password": "password123",
    "birthDate": "1990-01-01T00:00:00Z",
    "secretQuestion": "¿Cuál es tu color favorito?",
    "secretAnswer": "Azul"
  }'
```

### Iniciar Sesión
```bash
curl -X POST http://localhost:3000/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "juan@example.com",
    "password": "password123"
  }'
```

### Crear Tarea (requiere token JWT)
```bash
curl -X POST http://localhost:3000/api/v1/tasks/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <tu-jwt-token>" \
  -d '{
    "title": "Mi primera tarea",
    "description": "Descripción de la tarea",
    "startAt": "2024-01-01T09:00:00Z",
    "dueDate": "2024-01-01T17:00:00Z"
  }'
```

## 🏗️ Estructura del Proyecto


## 🔒 Autenticación

La API utiliza JWT (JSON Web Tokens) para la autenticación. Después del login exitoso, incluye el token en el header `Authorization` como `Bearer <token>` para acceder a las rutas protegidas.

## 🗄️ Modelos de Datos

### Usuario
```go
type User struct {
    ID             string    `json:"id,omitempty"`
    Name           string    `json:"name"`
    LastName       string    `json:"lastName"`
    Email          string    `json:"email"`
    Password       string    `json:"password"`
    BirthDate      time.Time `json:"birthDate"`
    SecretQuestion string    `json:"secretQuestion"`
    SecretAnswer   string    `json:"secretAnswer"`
}
```

### Tarea
```go
type Task struct {
    ID          string    `json:"id,omitempty"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    StartAt     time.Time `json:"startAt"`
    DueDate     time.Time `json:"dueDate"`
    UserID      string    `json:"userId"`
}
```

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## ✨ Autor

- Tu nombre - [@tu-usuario](https://github.com/tu-usuario)

## 🙏 Agradecimientos

- [Fiber](https://gofiber.io/) - Framework web para Go
- [MongoDB](https://www.mongodb.com/) - Base de datos NoSQL
- [JWT](https://jwt.io/) - Estándar para tokens de acceso
