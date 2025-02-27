package albums

import (
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dtoAlbum struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Cover       string               `bson:"cover"`
	ReleaseDate string               `bson:"release_date"`
	ArtistsIds  []primitive.ObjectID `bson:"artists_ids"`
	TracksIds   []primitive.ObjectID `bson:"tracks_ids"`
}
type dtoAlbumLight struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	ReleaseDate string             `bson:"release_date"`
	Cover       string             `bson:"cover"`
}

func (a *dtoAlbum) toCommonModel() models.AlbumPreliminary {
	return models.AlbumPreliminary{
		ID:          a.ID,
		Name:        a.Name,
		Cover:       a.Cover,
		ReleaseDate: a.ReleaseDate,
		ArtistsIds:  a.ArtistsIds,
		TracksIds:   a.TracksIds,
	}
}
func toCommonModels(albums []dtoAlbumLight) []models.AlbumLight {
	result := make([]models.AlbumLight, len(albums))

	for i, album := range albums {
		result[i] = models.AlbumLight{
			ID:          album.ID,
			Name:        album.Name,
			ReleaseDate: album.ReleaseDate,
		}
	}

	return result
}
