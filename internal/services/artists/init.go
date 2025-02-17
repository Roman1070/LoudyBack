package artists

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistsService struct {
	log     *slog.Logger
	artists Artists
}

type Artists interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
}

func New(log *slog.Logger, artists Artists) *ArtistsService {
	return &ArtistsService{
		artists: artists,
		log:     log,
	}
}
