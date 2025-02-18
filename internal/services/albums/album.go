package albums

import (
	"context"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	return s.albums.Album(ctx, id)
}
