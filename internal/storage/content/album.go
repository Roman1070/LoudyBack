package content

import (
	"context"
	models "loudy-back/internal/domain/models/content"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ContentStorage) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	return models.Album{}, nil
}
