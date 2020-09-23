package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func ExtractNotesWithFilter(client *mongo.Collection, filter bson.M) []*NoteResponse {
	cur, err := client.Find(context.TODO(), filter)
	var allNotesForUser []*NoteResponse
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem NoteResponse
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		allNotesForUser = append(allNotesForUser, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return allNotesForUser
}
