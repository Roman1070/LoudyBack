package albums

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/internal/services/artists"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AlbumsService struct {
	log     *slog.Logger
	albums  Albums
	artists artists.Artists
}

type Albums interface {
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
	CreateAlbum(ctx context.Context, name, cover string, releaseDate time.Time, artists_ids []primitive.ObjectID) (*emptypb.Empty, error)
}

func New(log *slog.Logger, albums Albums) *AlbumsService {
	return &AlbumsService{
		albums: albums,
		log:    log,
	}
}
