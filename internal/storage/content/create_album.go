package content

import (
	"context"
	models "loudy-back/internal/domain/models/content"
	"time"
)

func (s *ContentStorage) CreateAlbum(ctx context.Context, name, cover string, tracks_ids []models.TrackLight, releaseDate time.Time) (uint32, error) {

	return 0, nil
}
