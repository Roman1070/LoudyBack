package profiles

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) ToggleLikeAlbum(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID) (bool, error) {
	s.log.Info("[ToggleLikeAlbum] service started")

	liked, err := s.profiles.IsAlbumLiked(ctx, profileId, albumId)
	if err != nil {
		s.log.Error("[ToggleLikeAlbum] service error: " + err.Error())
		return false, errors.New("[ToggleLikeAlbum] service error: " + err.Error())
	}

	err = s.profiles.SetAlbumLiked(ctx, profileId, albumId, !liked)
	if err != nil {
		s.log.Error("[ToggleLikeAlbum] service error: " + err.Error())
		return false, errors.New("[ToggleLikeAlbum] service error: " + err.Error())
	}

	return !liked, nil
}
