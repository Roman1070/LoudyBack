package tracks

import (
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dtoTrack struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Name       string               `bson:"name"`
	Filename   string               `bson:"filename"`
	AlbumId    primitive.ObjectID   `bson:"album_id"`
	ArtistsIds []primitive.ObjectID `bson:"artists_ids"`
	Duration   uint16               `bson:"duration"`
}

type dtoTrackLight struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Filename string             `bson:"filename"`
	AlbumId  primitive.ObjectID `bson:"album_id"`
	Duration uint16             `bson:"duration"`
}

func toLightModels(tracks []dtoTrackLight) []models.TrackLight {
	result := make([]models.TrackLight, len(tracks))

	for i, track := range tracks {
		result[i] = models.TrackLight{
			ID:       track.ID,
			Name:     track.Name,
			Filename: track.Filename,
			Duration: track.Duration,
		}
	}

	return result
}
