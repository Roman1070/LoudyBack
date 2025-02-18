package models

import (
	artistsv1 "loudy-back/gen/go/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArtistLight struct {
	ID    primitive.ObjectID
	Name  string
	Cover string
}

func (artist *ArtistLight) ToGRPC() *artistsv1.ArtistLight {

	return &artistsv1.ArtistLight{
		Id:   artist.ID.Hex(),
		Name: artist.Name,
	}
}
