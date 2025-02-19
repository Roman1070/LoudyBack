package artists

import (
	"context"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistsService struct {
	log             *slog.Logger
	artists         Artists
	albums          albumsv1.AlbumsClient
	artistsProvider ArtistsProvider
}

type Artists interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
	AddAlbum(ctx context.Context, artistsIds []primitive.ObjectID, albumId primitive.ObjectID) (*emptypb.Empty, error)
}
type ArtistsProvider interface {
	ArtistLightByName(ctx context.Context, name string) (models.ArtistLight, error)
}

func New(log *slog.Logger, artists Artists, albums albumsv1.AlbumsClient, artistsProvider ArtistsProvider) *ArtistsService {
	return &ArtistsService{
		artists:         artists,
		albums:          albums,
		log:             log,
		artistsProvider: artistsProvider,
	}
}
