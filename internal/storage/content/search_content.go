package content

import (
	"context"
	models "loudy-back/internal/domain/models/content"
)

func (s *ContentStorage) SearchContent(ctx context.Context, input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error) {
	return nil, nil, nil, nil
}
