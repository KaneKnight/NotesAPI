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

	//Connect()
	
	r.POST("/login", gah.LoginHandler)
    r.POST("/register", gah.RegisterHandler)
	r.Run(":8080")
}



func Connect() {
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

	collection := client.Database("thirdfort").Collection("users")

	_, err = collection.InsertOne(context.TODO(), gah.LoginStruct{"kaneywaney11@hotmail.co.uk", "12345"})
	if err != nil {
		log.Fatal(err)
	}

	// Close Connection
	err = client.Disconnect(context.TODO())
	if err != nil {
   	 log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}