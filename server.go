package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasaricli/gah"
	"context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  )

func main() {
	r := gin.Default()

	db := Connect()

	r.Use(DbMiddleware(db))
	
	r.POST("/login", gah.LoginHandler)
	r.POST("/register", gah.RegisterHandler)
	r.POST("/notes", gah.AuthRequiredMiddleware, CreateNoteHandler)
	r.PUT("/notes/:id", gah.AuthRequiredMiddleware, UpdateNoteHandler)
	r.DELETE("/notes/:id", gah.AuthRequiredMiddleware, DeleteNoteHandler)

	r.PUT("/notes/:id/archive", gah.AuthRequiredMiddleware, ArchiveNoteHandler)

	r.GET("/notes", gah.AuthRequiredMiddleware, GetNoteHandler)

	r.Run(":8080")
}

// Middleware will add the db connection to the context
func DbMiddleware(db* mongo.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("client", db)
        c.Next()
    }
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}