package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID          primitive.ObjectID
	Artists     []ArtistLight
	Name        string
	Cover       string
	Tracks      []TrackLight
	ReleaseDate time.Time
}
