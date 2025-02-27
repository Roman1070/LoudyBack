package albums

import (
	"context"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	tracksModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArtistsProvider interface {
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]artistsModels.ArtistLight, error)
}
type TracksProvider interface {
	TracksLight(ctx context.Context, ids []primitive.ObjectID) ([]tracksModels.TrackLight, error)
}

type AlbumsService struct {
	log             *slog.Logger
	albums          Albums
	artists         artistsv1.ArtistsClient
	artistsProvider ArtistsProvider
	tracksProvider  TracksProvider
}

type Albums interface {
	CreateAlbum(ctx context.Context, name, cover string,
		releaseDate string, artists_ids []primitive.ObjectID) (primitive.ObjectID, error)
	Album(ctx context.Context, id primitive.ObjectID) (models.AlbumPreliminary, error)
}

func New(log *slog.Logger, artists artistsv1.ArtistsClient, albums Albums, artistsProvider ArtistsProvider, tracksProvider TracksProvider) *AlbumsService {
	return &AlbumsService{
		albums:          albums,
		log:             log,
		artistsProvider: artistsProvider,
		artists:         artists,
		tracksProvider:  tracksProvider,
	}
}
