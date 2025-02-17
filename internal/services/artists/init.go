package artists

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistsService struct {
	log             *slog.Logger
	artistsCreator  ArtistsCreator
	artistsProvider ArtistsProvider
}

type ArtistsProvider interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
}

type ArtistsCreator interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
}

func New(log *slog.Logger, artistsCreator ArtistsCreator, artistsProvider ArtistsProvider) *ArtistsService {
	return &ArtistsService{
		artistsCreator:  artistsCreator,
		artistsProvider: artistsProvider,
		log:             log,
	}
}
