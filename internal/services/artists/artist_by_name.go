package artists

import (
	"context"
	models "loudy-back/internal/domain/models/artists"
)

func (s *ArtistsService) ArtistByName(ctx context.Context, name string) (models.Artist, error) {
	return s.artists.ArtistByName(ctx, name)
}
