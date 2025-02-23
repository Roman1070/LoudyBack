package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Track struct {
	ID       primitive.ObjectID
	Name     string
	AlbumID  primitive.ObjectID
	Cover    string
	Artists  []ArtistLight
	Duration uint16
}

type ArtistLight struct {
	ID   primitive.ObjectID
	Name string
}
