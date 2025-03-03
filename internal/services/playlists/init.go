package playlists

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/playlists"
	trackModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaylistsService struct {
	log            *slog.Logger
	playlists      Playlists
	tracksProvider TracksProvider
}

type TracksProvider interface {
	Tracks(ctx context.Context, ids []primitive.ObjectID) ([]trackModels.Track, error)
}

type Playlists interface {
	Playlist(ctx context.Context, id primitive.ObjectID) (models.PlaylistPreliminary, error)
	PlaylistLight(ctx context.Context, id primitive.ObjectID) (models.PlaylistLight, error)
}

func New(log *slog.Logger, playlists Playlists, tracksProvider TracksProvider) *PlaylistsService {
	return &PlaylistsService{
		log:            log,
		playlists:      playlists,
		tracksProvider: tracksProvider,
	}
}
