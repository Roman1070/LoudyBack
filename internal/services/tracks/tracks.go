package tracks

import (
	"context"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *TracksService) Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.Track, error) {
	return nil, nil
}
