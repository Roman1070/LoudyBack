package artists

import (
	"context"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ArtistsService) Artist(ctx context.Context, id primitive.ObjectID) (models.Artist, error) {
	return s.artists.Artist(ctx, id)
}
