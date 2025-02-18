package albums

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *AlbumsStorage) AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error) {
	s.log.Info("[AlbumsLight] storage started")

	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.collection.Find(ctx, query)
	if err != nil {
		s.log.Info("[AlbumsLight] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrAlbumNotFound
		}

		slog.Error("[AlbumsLight] storage error: " + err.Error())
		return nil, errors.New("[AlbumsLight] storage error: " + err.Error())
	}

	var results []dtoAlbumLight
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[AlbumsLight] storage error: " + err.Error())
		return nil, errors.New("[AlbumsLight] storage error: " + err.Error())
	}

	return toCommonModels(results), nil
}
