package albums

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/albums"
	trackModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	s.log.Info("[Album] service started")

	preliminaryAlbum, err := s.albums.Album(ctx, id)
	if err != nil {
		s.log.Error("[Album] service error: " + err.Error())
		return models.Album{}, errors.New("[Album] service error: " + err.Error())
	}

	artists, err := s.artistsProvider.ArtistsLight(ctx, preliminaryAlbum.ArtistsIds)
	if err != nil {
		s.log.Error("[Album] service error: " + err.Error())
		return models.Album{}, errors.New("[Album] service error: " + err.Error())
	}

	tracks := []trackModels.TrackLight{}

	return models.Album{
		ID:          preliminaryAlbum.ID,
		Name:        preliminaryAlbum.Name,
		Cover:       preliminaryAlbum.Cover,
		ReleaseDate: preliminaryAlbum.ReleaseDate,
		Artists:     artists,
		Tracks:      tracks,
	}, nil
}
