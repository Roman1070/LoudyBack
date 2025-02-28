package tracks

import (
	"context"
	"log/slog"
	albumModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TracksService struct {
	log             *slog.Logger
	tracks          Tracks
	artistsProvider ArtistsProvider
	albumsProvider  AlbumsProvider
}

type ArtistsProvider interface {
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]artistsModels.ArtistLight, error)
}
type AlbumsProvider interface {
	AlbumLight(ctx context.Context, id primitive.ObjectID) (albumModels.AlbumLight, error)
}

type Tracks interface {
	CreateTrack(ctx context.Context, name, filename string, albumId primitive.ObjectID, artistsIds []primitive.ObjectID, duration uint16) (primitive.ObjectID, error)
	Track(ctx context.Context, id primitive.ObjectID) (models.TrackPreliminary, error)
	Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.TrackPreliminary, error)
}

func New(tracks Tracks, artistsProvider ArtistsProvider, albumsProvider AlbumsProvider, log *slog.Logger) *TracksService {
	return &TracksService{
		tracks:          tracks,
		artistsProvider: artistsProvider,
		albumsProvider:  albumsProvider,
		log:             log,
	}
}
