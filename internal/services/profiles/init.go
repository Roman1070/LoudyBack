package profiles

import (
	"context"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	albumsModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	models "loudy-back/internal/domain/models/profiles"
	tracksModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfilesService struct {
	log             *slog.Logger
	profiles        Profiles
	artists         artistsv1.ArtistsClient
	artistsProvider ArtistsProvider
	albumsProvider  AlbumsProvider
	tracksProvider  TracksProvider
}

type ArtistsProvider interface {
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]artistsModels.ArtistLight, error)
}

type AlbumsProvider interface {
	AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]albumsModels.AlbumLight, error)
}

type TracksProvider interface {
	Tracks(ctx context.Context, ids []primitive.ObjectID) ([]tracksModels.Track, error)
}

type Profiles interface {
	CreateProfile(ctx context.Context, userId uint32, name, avatar, bio string) (id primitive.ObjectID, err error)
	Profile(ctx context.Context, id primitive.ObjectID) (profile models.ProfilePreliminary, err error)
}

func New(log *slog.Logger, artists artistsv1.ArtistsClient, profiles Profiles, artistsProvider ArtistsProvider, albumsProvider AlbumsProvider, tracksProvider TracksProvider) *ProfilesService {
	return &ProfilesService{
		profiles:        profiles,
		log:             log,
		artistsProvider: artistsProvider,
		albumsProvider:  albumsProvider,
		tracksProvider:  tracksProvider,
		artists:         artists,
	}
}
