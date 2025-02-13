package content

import (
	"context"
	"time"
)

func (s *ContentService) CreateAlbum(ctx context.Context, name, cover string, tracksIds []uint32, releaseDate time.Time) (uint32, error) {
	return 0, nil
}
