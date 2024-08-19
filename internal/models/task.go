package models

import "time"

type Task struct {
	ID        int       `json:"id" bson:"id"`
	Task      string    `json:"task" bson:"task"`
	StartTime time.Time `json:"startTime" bson:"startTime"`
	EndTime   time.Time `json:"endTime" bson:"endTime"`
	Color     string    `json:"color" bson:"color"`
	UserID    string    `json:"userId" bson:"userId"`
}
