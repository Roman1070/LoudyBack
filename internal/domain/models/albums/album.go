package models

import (
	models "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	ReleaseDate time.Time
	Artists     []models.Artist
	Tracks      []trackModels.TrackLight
}
