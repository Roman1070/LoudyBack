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

func (s *TracksStorage) Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.TrackPreliminary, error) {
	s.log.Info("[Tracks] storage started")

	if len(ids) == 0 {
		return []models.TrackPreliminary{}, nil
	}

	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.collection.Find(ctx, query)
	if err != nil {
		s.log.Info("[Tracks] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrTrackNotFound
		}

		slog.Error("[Tracks] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Tracks] storage error: "+err.Error())
	}
	s.log.Info("[Tracks] cursor recieved")

	var results []dtoTrack
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[Tracks] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Tracks] storage error: "+err.Error())
	}

	return toCommonModels(results), nil
}
