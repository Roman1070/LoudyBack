package profiles

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/profiles"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) Profile(ctx context.Context, id primitive.ObjectID) (models.Profile, error) {
	s.log.Info("[GetProfile] server started")

	profile, err := s.profiles.Profile(ctx, id)
	if err != nil {
		s.log.Error("[GetProfile] service error: " + err.Error())
		return models.Profile{}, errors.New("[GetProfile] service error: " + err.Error())
	}

	artists, err := s.artistsProvider.ArtistsLight(ctx, profile.SavedArtistsIds)
	if err != nil {
		s.log.Error("[GetProfile] service error: " + err.Error())
		return models.Profile{}, errors.New("[GetProfile] service error: " + err.Error())
	}

	albums, err := s.albumsProvider.AlbumsLight(ctx, profile.SavedAlbumsIds)
	if err != nil {
		s.log.Error("[GetProfile] service error: " + err.Error())
		return models.Profile{}, errors.New("[GetProfile] service error: " + err.Error())
	}

	tracks, err := s.tracksProvider.Tracks(ctx, profile.SavedTracksIds)
	if err != nil {
		s.log.Error("[GetProfile] service error: " + err.Error())
		return models.Profile{}, errors.New("[GetProfile] service error: " + err.Error())
	}

	return models.Profile{
		ID:           profile.ID,
		Name:         profile.Name,
		Avatar:       profile.Avatar,
		Bio:          profile.Bio,
		SavedTracks:  tracks,
		SavedAlbums:  albums,
		SavedArtists: artists,
	}, nil
}
