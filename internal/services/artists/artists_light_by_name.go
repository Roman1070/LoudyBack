package artists

import (
	"context"
	models "loudy-back/internal/domain/models/artists"
)

func (s *ArtistsService) ArtistLightByName(ctx context.Context, name string) (models.ArtistLight, error) {
	return s.artists.ArtistLightByName(ctx, name)
}
