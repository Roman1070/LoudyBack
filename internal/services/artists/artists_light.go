package artists

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ArtistsService) ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error) {
	s.log.Info("[Artists] service started")

	artists, err := s.artists.ArtistsLight(ctx, ids)
	if err != nil {
		s.log.Error("[Artists] service error: " + err.Error())
		return nil, errors.New("[Artists] service error: " + err.Error())
	}

	return artists, nil
}
