package profiles

import (
	"context"
	"errors"
	"log/slog"
	"slices"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesStorage) IsTrackLiked(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID) (bool, error) {
	s.log.Info("[IsTrackLiked] storage started")

	profile, err := s.Profile(ctx, profileId)
	if err != nil {
		slog.Error("[IsTrackLiked] storage error: " + err.Error())
		return false, errors.New("[IsTrackLiked] storage error: " + err.Error())
	}

	return slices.Contains(profile.SavedTracksIds, trackId), nil
}
