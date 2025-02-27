package tracks

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/tracks"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TracksStorage) Track(ctx context.Context, id primitive.ObjectID) (models.TrackPreliminary, error) {
	s.log.Info("[Track] storage started")

	filter := bson.M{"_id": id}

	var track dtoTrack

	err := s.collection.FindOne(ctx, filter).Decode(&track)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.TrackPreliminary{}, storage.ErrTrackNotFound
		}

		slog.Error("[Profile] storage error: " + err.Error())
		return models.TrackPreliminary{}, errors.New("[Profile] storage error: " + err.Error())
	}

	return models.TrackPreliminary{
		ID:         track.ID,
		Name:       track.Name,
		AlbumID:    track.AlbumId,
		ArtistsIds: track.ArtistsIds,
		Duration:   track.Duration,
	}, nil
}
