package content

import (
	"context"
	models "loudy-back/internal/domain/models/content"
)

func (s *ContentStorage) Album(ctx context.Context, id uint32) (models.Album, error) {
	return models.Album{}, nil
}
