package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Content string `json:"content" binding:"required"`
	Archived bool `json:"archived"`
	UserID primitive.ObjectID `json:"userID"`
}