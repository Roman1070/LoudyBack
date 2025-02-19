package artists

import (
	"context"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistsService struct {
	log     *slog.Logger
	artists Artists
	albums  albumsv1.AlbumsClient
}

type Artists interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
	AddAlbum(ctx context.Context, artistsIds []primitive.ObjectID, albumId primitive.ObjectID) (*emptypb.Empty, error)
}

func New(log *slog.Logger, artists Artists, albums albumsv1.AlbumsClient) *ArtistsService {
	return &ArtistsService{
		artists: artists,
		albums:  albums,
		log:     log,
	}
}
