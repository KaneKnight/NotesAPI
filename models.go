package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Content  string             `json:"content" binding:"required"`
	Archived bool               `json:"archived"`
	UserID   primitive.ObjectID `json:"userID"`
}

type NoteResponse struct {
	Content  string             `json:"content" binding:"required"`
	Archived bool               `json:"archived"`
	UserID   primitive.ObjectID `json:"userID"`
	ID       primitive.ObjectID `bson:"_id"`
}

type ArchivePayload struct {
	Archived bool `json:"archived"`
}
