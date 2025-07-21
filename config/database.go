package config

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client // Cliente de MongoDB
var ctx context.Context  // Contexto para operaciones asíncronas

// ConnectToDatabase establece la conexión con MongoDB
func ConnectToDatabase() {
    // Leer la URI de conexión desde una variable de entorno
    mongodbURI := os.Getenv("MONGODB_URI")
    if mongodbURI == "" {
        log.Fatal("La variable de entorno MONGODB_URI no está configurada")
    }

    // Configurar el contexto con un timeout de 10 segundos
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Configurar las opciones del cliente
    clientOptions := options.Client().ApplyURI(mongodbURI)

    // Conectar al cliente
    var err error
    Client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Error al conectar a MongoDB: %v", err)
    }

    // Verificar la conexión
    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("No se pudo pingear MongoDB: %v", err)
    }

    log.Println("Conexión exitosa a MongoDB")
}

// GetCollection devuelve una colección específica de la base de datos
func GetCollection(databaseName, collectionName string) *mongo.Collection {
    return Client.Database(databaseName).Collection(collectionName)
}