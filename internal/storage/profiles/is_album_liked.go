package profiles

import (
	"context"
	"errors"
	"log/slog"
	"slices"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesStorage) IsAlbumLiked(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID) (bool, error) {
	s.log.Info("[IsAlbumLiked] storage started")

	profile, err := s.Profile(ctx, profileId)
	if err != nil {
		slog.Error("[IsAlbumLiked] storage error: " + err.Error())
		return false, errors.New("[IsAlbumLiked] storage error: " + err.Error())
	}

	return slices.Contains(profile.SavedAlbumsIds, albumId), nil
}
