package tracks

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *TracksService) CreateTrack(ctx context.Context, name, filename string, albumId primitive.ObjectID, artistsIds []primitive.ObjectID, duration uint16) (primitive.ObjectID, error) {
	return s.tracks.CreateTrack(ctx, name, filename, albumId, artistsIds, duration)
}
