package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrackLight struct {
	ID       primitive.ObjectID `bson:"omitempty,_id"`
	Name     string
	Duration uint16
}
