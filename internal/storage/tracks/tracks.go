package tracks

import (
	"context"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *TracksStorage) Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.TrackPreliminary, error) {
	return nil, nil
}
