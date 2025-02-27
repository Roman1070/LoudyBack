package tracks

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/tracks"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TracksStorage) TracksLight(ctx context.Context, ids []primitive.ObjectID) ([]models.TrackLight, error) {
	s.log.Info("[TracksLight] storage started")

	if len(ids) == 0 {
		return []models.TrackLight{}, nil
	}

	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.collection.Find(ctx, query)
	if err != nil {
		s.log.Info("[TracksLight] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrTrackNotFound
		}

		slog.Error("[TracksLight] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[TracksLight] storage error: "+err.Error())
	}
	s.log.Info("[Artists] cursor recieved")

	var results []dtoTrackLight
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[Artists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artists] storage error: "+err.Error())
	}

	return toLightModels(results), nil
}
