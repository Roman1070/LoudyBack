package profiles

import (
	"context"
	"errors"
	"log/slog"
	"slices"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesStorage) IsArtistLiked(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID) (bool, error) {
	s.log.Info("[IsArtistLiked] storage started")

	profile, err := s.Profile(ctx, profileId)
	if err != nil {
		slog.Error("[IsArtistLiked] storage error: " + err.Error())
		return false, errors.New("[IsArtistLiked] storage error: " + err.Error())
	}

	return slices.Contains(profile.SavedArtistsIds, artistId), nil
}
