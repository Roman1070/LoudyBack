package profiles

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) ToggleLikeTrack(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID) (bool, error) {
	s.log.Info("[ToggleLikeTrack] service started")

	liked, err := s.profiles.IsTrackLiked(ctx, profileId, trackId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] service error: " + err.Error())
		return false, errors.New("[ToggleLikeTrack] service error: " + err.Error())
	}

	err = s.profiles.SetTrackLiked(ctx, profileId, trackId, !liked)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] service error: " + err.Error())
		return false, errors.New("[ToggleLikeTrack] service error: " + err.Error())
	}

	return !liked, nil
}
