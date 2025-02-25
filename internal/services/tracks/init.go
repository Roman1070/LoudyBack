package tracks

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TracksService struct {
	log             *slog.Logger
	tracks          Tracks
	artistsProvider ArtistsProvider
}

type ArtistsProvider interface {
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error)
}

type Tracks interface {
	CreateTrack(ctx context.Context, name string, albumId primitive.ObjectID, artistsIds []primitive.ObjectID, duration uint16) (primitive.ObjectID, error)
	Track(ctx context.Context, id primitive.ObjectID) (models.TrackPreliminary, error)
	Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.TrackPreliminary, error)
}

func New(tracks Tracks, artistsProvider ArtistsProvider, log *slog.Logger) *TracksService {
	return &TracksService{
		tracks:          tracks,
		artistsProvider: artistsProvider,
		log:             log,
	}
}
