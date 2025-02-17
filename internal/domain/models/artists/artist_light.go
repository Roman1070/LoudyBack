package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ArtistLight struct {
	ID    primitive.ObjectID
	Name  string
	Cover string
}
