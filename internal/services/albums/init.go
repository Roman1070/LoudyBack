package albums

import (
	"context"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AlbumsService struct {
	log     *slog.Logger
	albums  Albums
	artists artistsv1.ArtistsClient
}

type Albums interface {
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
	AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error)
	CreateAlbum(ctx context.Context, name, cover string, releaseDate string, artists_ids []primitive.ObjectID) (*emptypb.Empty, error)
}

func New(log *slog.Logger, artists artistsv1.ArtistsClient, albums Albums) *AlbumsService {
	return &AlbumsService{
		albums:  albums,
		log:     log,
		artists: artists,
	}
}
