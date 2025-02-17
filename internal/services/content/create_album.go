package content

import (
	"context"
	models "loudy-back/internal/domain/models/content"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ContentService) CreateAlbum(ctx context.Context, name, cover string, tracks []models.TrackLight, releaseDate time.Time) (*emptypb.Empty, error) {
	return nil, nil
}
