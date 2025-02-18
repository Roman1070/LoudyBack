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
	log     *slog.Logger
	artists Artists
	albums  albumsv1.AlbumsClient
}

type Artists interface {
	Artist(ctx context.Context, id string) (models.Artist, error)
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error)
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
