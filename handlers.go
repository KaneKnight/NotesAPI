package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"context"
	"net/http"
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

	res, err := client.Database("thirdfort").Collection("notes").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if res.MatchedCount == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,  gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Note does not exist so I can't update it.",
			},
		})
	} else {
		c.JSON(200, gin.H{
			"status"  : "Note updated.",
			"content" : note.Content,
			"noteID"  : noteID,
			"userID"  : note.UserID})
	}
}

func DeleteNoteHandler(c* gin.Context) {
	client := c.MustGet("client").(*mongo.Client)
	noteID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": bson.M{"$eq": noteID}}

	res, err := client.Database("thirdfort").Collection("notes").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if res.DeletedCount == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,  gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Note does not exist so I can't delete it.",
			},
		})	
	} else {
		c.JSON(200, gin.H{
		"status"  : "Note deleted."})
	}	
}

func ArchiveNoteHandler(c* gin.Context) {
	
}

func GetNoteHandler(c* gin.Context) {
}