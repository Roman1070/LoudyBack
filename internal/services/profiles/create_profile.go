package profiles

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) CreateProfile(ctx context.Context, userId uint32, name, avatar, bio string) (id primitive.ObjectID, err error) {
	return s.profiles.CreateProfile(ctx, userId, name, avatar, bio)
}
