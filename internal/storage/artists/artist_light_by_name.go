package artists

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ArtistsStorage) ArtistLightByName(ctx context.Context, name string) (models.ArtistLight, error) {
	s.log.Info("[ArtistLightByName] storage started")

	filter := bson.M{"name": name}

	var result dtoArtistLight

	err := s.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.ArtistLight{}, storage.ErrArtistNotFound
		}

		slog.Error("[ArtistLightByName] storage error: " + err.Error())
		return models.ArtistLight{}, errors.New("[ArtistLightByName] storage error: " + err.Error())
	}

	return models.ArtistLight{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}
