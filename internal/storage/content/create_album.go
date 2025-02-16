package content

import (
	"context"
	"time"
)

func (s *ContentStorage) CreateAlbum(ctx context.Context, name, cover string, tracks_ids []uint32, releaseDate time.Time) (uint32, error) {
	return 0, nil
}
