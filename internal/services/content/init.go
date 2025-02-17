package content

import (
	"context"
	"log/slog"

	models "loudy-back/internal/domain/models/content"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentService struct {
	log             *slog.Logger
	contentCreator  ContentCreator
	contentProvider ContentProvider
}

type ContentProvider interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
	SearchContent(ctx context.Context, input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error)
}

type ContentCreator interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
	CreateTrack(ctx context.Context, name, file string, albumId primitive.ObjectID) (*emptypb.Empty, error)
	CreateAlbum(ctx context.Context, name, cover string, tracksIds []models.TrackLight, releaseDate time.Time) (uint32, error)
}

func New(log *slog.Logger, contentCreator ContentCreator, contentProvider ContentProvider) *ContentService {
	return &ContentService{
		contentCreator:  contentCreator,
		contentProvider: contentProvider,
		log:             log,
	}
}
