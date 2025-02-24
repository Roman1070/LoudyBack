package profiles

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/profiles"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *ProfilesStorage) Profile(ctx context.Context, id primitive.ObjectID) (models.ProfilePreliminary, error) {
	s.log.Info("[Profile] storage started")

	filter := bson.M{"_id": id}

	var profile dtoProfile

	err := s.collection.FindOne(ctx, filter, &options.FindOneOptions{}).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.ProfilePreliminary{}, storage.ErrProfileNotFound
		}

		slog.Error("[Profile] storage error: " + err.Error())
		return models.ProfilePreliminary{}, errors.New("[Profile] storage error: " + err.Error())
	}
	return models.ProfilePreliminary{
		ID:                profile.ID,
		Name:              profile.Name,
		Avatar:            profile.Avatar,
		Bio:               profile.Bio,
		LikesCount:        profile.LikesCount,
		UserId:            profile.UserId,
		SavedAlbumsIds:    profile.SavedAlbumsIds,
		SavedTracksIds:    profile.SavedTracksIds,
		SavedArtistsIds:   profile.SavedArtistsIds,
		SavedPlaylistsIds: profile.SavedPlaylistsIds,
	}, nil
}
