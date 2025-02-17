package albums

import (
	models "loudy-back/internal/domain/models/albums"
	artistModels "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dtoAlbum struct {
	ID          primitive.ObjectID   `bson:"omitempty,_id"`
	Name        string               `bson:"name"`
	Cover       string               `bson:"cover"`
	ReleaseDate time.Time            `bson:"release_date"`
	ArtistsIds  []primitive.ObjectID `bson:"artists_ids"`
	TracksIds   []primitive.ObjectID `bson:"tracks_ids"`
}

func (a *dtoAlbum) toCommonModel(artists []artistModels.Artist, tracks []trackModels.TrackLight) models.Album {
	return models.Album{
		ID:          a.ID,
		Name:        a.Name,
		Cover:       a.Cover,
		ReleaseDate: a.ReleaseDate,
		Artists:     artists,
		Tracks:      tracks,
	}
}
