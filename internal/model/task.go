package model

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    UserID      string             `json:"userId" bson:"userId"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    Completed   bool               `json:"completed" bson:"completed"`
}