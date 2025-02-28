package profiles

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) ToggleLikeArtist(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID) (bool, error) {
	s.log.Info("[ToggleLikeArtist] service started")

	liked, err := s.profiles.IsArtistLiked(ctx, profileId, artistId)
	if err != nil {
		s.log.Error("[ToggleLikeArtist] service error: " + err.Error())
		return false, errors.New("[ToggleLikeArtist] service error: " + err.Error())
	}

	err = s.profiles.SetArtistLiked(ctx, profileId, artistId, !liked)
	if err != nil {
		s.log.Error("[ToggleLikeArtist] service error: " + err.Error())
		return false, errors.New("[ToggleLikeArtist] service error: " + err.Error())
	}

	return !liked, nil
}
