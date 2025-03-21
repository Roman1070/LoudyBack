package profiles

import (
	"context"
	"log/slog"
	albumsModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	models "loudy-back/internal/domain/models/profiles"
	tracksModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfilesService struct {
	log             *slog.Logger
	profiles        Profiles
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
	IsTrackLiked(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID) (bool, error)
	SetTrackLiked(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID, liked bool) error
	IsAlbumLiked(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID) (bool, error)
	SetAlbumLiked(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID, liked bool) error
	IsArtistLiked(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID) (bool, error)
	SetArtistLiked(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID, liked bool) error
}

func New(log *slog.Logger, profiles Profiles, artistsProvider ArtistsProvider, albumsProvider AlbumsProvider, tracksProvider TracksProvider) *ProfilesService {
	return &ProfilesService{
		profiles:        profiles,
		log:             log,
		artistsProvider: artistsProvider,
		albumsProvider:  albumsProvider,
		tracksProvider:  tracksProvider,
	}
}
