package albums

import (
	models "loudy-back/internal/domain/models/albums"
	artistModels "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"

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
}

func (a *dtoAlbum) toCommonModel(artists []artistModels.ArtistLight, tracks []trackModels.TrackLight) models.Album {
	return models.Album{
		ID:          a.ID,
		Name:        a.Name,
		Cover:       a.Cover,
		ReleaseDate: a.ReleaseDate,
		Artists:     artists,
		Tracks:      tracks,
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
