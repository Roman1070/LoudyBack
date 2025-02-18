package albums

import (
	"context"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error) {
	return s.albums.AlbumsLight(ctx, ids)
}
