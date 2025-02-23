package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrackLight struct {
	ID       primitive.ObjectID
	Name     string
	Duration uint16
}
