package models

import (
	models "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	ReleaseDate string
	Artists     []models.ArtistLight
	Tracks      []trackModels.TrackLight
}
