package profiles

import (
	"context"
	"errors"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesStorage) CreateProfile(ctx context.Context, userId uint32, name, avatar, bio string) (primitive.ObjectID, error) {
	s.log.Info("[CreateProfile] storage started")

	profile := dtoProfile{
		ID:         primitive.NewObjectID(),
		UserId:     userId,
		Name:       name,
		Avatar:     avatar,
		Bio:        bio,
		LikesCount: 0,
	}

	_, err := s.collection.InsertOne(ctx, profile)
	if err != nil {
		slog.Error("[CreateProfile] storage error: " + err.Error())
		return [12]byte{}, errors.New("[CreateProfile] storage error: " + err.Error())
	}

	return profile.ID, nil
}
