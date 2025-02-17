package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlbumLight struct {
	ID    primitive.ObjectID
	Name  string
	Cover string
	Year  uint32
}
