package artists

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/artists"
)

func (s *ArtistsService) Artist(ctx context.Context, name string) (models.Artist, error) {
	s.log.Info("[Artist] service started")

	artist, err := s.artists.Artist(ctx, name)
	if err != nil {
		s.log.Error("[Artist] service error: " + err.Error())
		return models.Artist{}, errors.New("[Artist] service error: " + err.Error())
	}

	return artist, nil
}
