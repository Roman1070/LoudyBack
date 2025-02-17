package albums

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	s.log.Info("[Album] service started")

	album, err := s.albums.Album(ctx, id)
	if err != nil {
		s.log.Error("[Album] service error: " + err.Error())
		return models.Album{}, errors.New("[Album] service error: " + err.Error())
	}

	return album, nil
}
