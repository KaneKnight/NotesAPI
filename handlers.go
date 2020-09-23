package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func CreateNoteHandler(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	client := c.MustGet("client").(*mongo.Collection)

	var note Note
	c.BindJSON(&note)
	note.Archived = false
	note.UserID = userID

	res, err := client.InsertOne(context.TODO(), note)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"status":  "Note saved.",
		"content": note.Content,
		"noteID":  res.InsertedID,
		"userID":  note.UserID})
}

func UpdateNoteHandler(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	client := c.MustGet("client").(*mongo.Collection)
	noteID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var note Note
	c.BindJSON(&note)
	note.UserID = userID

	filter := bson.M{"_id": bson.M{"$eq": noteID},
		"archived": bson.M{"$eq": false}}
	update := bson.M{"$set": bson.M{"content": note.Content}}

	res, err := client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	} else if res.MatchedCount == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Note does not exist or is archived so I can't update it.",
			},
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "Note updated.",
			"content": note.Content,
			"noteID":  noteID,
			"userID":  note.UserID})
	}
}

func DeleteNoteHandler(c *gin.Context) {
	client := c.MustGet("client").(*mongo.Collection)
	noteID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": bson.M{"$eq": noteID},
		"archived": bson.M{"$eq": false}}

	res, err := client.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	} else if res.DeletedCount == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Note does not exist so I can't delete it.",
			},
		})
	} else {
		c.JSON(200, gin.H{
			"status": "Note deleted."})
	}
}

func ArchiveNoteHandler(c *gin.Context) {
	client := c.MustGet("client").(*mongo.Collection)
	noteID, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var payload ArchivePayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Payload is not correct.",
			},
		})
		return
	}

	filter := bson.M{"_id": bson.M{"$eq": noteID}}
	update := bson.M{"$set": bson.M{"archived": payload.Archived}}

	res, err := client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	} else if res.MatchedCount == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Note does not exist so I can't change its archive status.",
			},
		})
	} else {
		if payload.Archived {
			c.JSON(200, gin.H{
				"status": "Note archived."})
		} else {
			c.JSON(200, gin.H{
				"status": "Note unarchived."})
		}
	}

}

func GetNoteHandler(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)
	client := c.MustGet("client").(*mongo.Collection)
	archivedStr := c.Query("archived")
	var filter bson.M
	archived, err := strconv.ParseBool(archivedStr)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"data": gin.H{
				"message": "Query string should be of form archived={true|false}",
			},
		})
	} else {
		filter = bson.M{"userid": bson.M{"$eq": userID},
			"archived": bson.M{"$eq": archived}}

		allNotesForUser := ExtractNotesWithFilter(client, filter)
		if allNotesForUser == nil {
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{
				"status": "error",
				"data": gin.H{
					"message": "Query was well formed but there are no notes.",
				},
			})
		} else {
			c.JSON(200, gin.H{
				"status": "Notes collected.",
				"notes":  allNotesForUser})
		}
	}
}
