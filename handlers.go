package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"context"
)

func CreateNoteHandler(c* gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	client := c.MustGet("client").(*mongo.Client)

	var note Note
	c.BindJSON(&note)
	note.Archived = false
	note.UserID = userID

	res, err := client.Database("thirdfort").Collection("notes").InsertOne(context.TODO(), note)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"status"  : "Note saved.",
		"content" : note.Content,
		"noteID"  : res.InsertedID,
		"userID"  : note.UserID})
}

func UpdateNoteHandler(c* gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	client := c.MustGet("client").(*mongo.Client)
	noteID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var note Note
	c.BindJSON(&note)
	note.UserID = userID
	
	filter := bson.M{"_id": bson.M{"$eq": noteID}}
	update := bson.M{"$set": bson.M{"content": note.Content}}

	_, err := client.Database("thirdfort").Collection("notes").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"status"  : "Note updated.",
		"content" : note.Content,
		"noteID"  : noteID,
		"userID"  : note.UserID})
}

func DeleteNoteHandler(c* gin.Context) {
	
}

func ArchiveNoteHandler(c* gin.Context) {
	
}

func GetNoteHandler(c* gin.Context) {
}