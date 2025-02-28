package profiles

import (
	"context"
	"errors"
	"log/slog"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *ProfilesStorage) SetTrackLiked(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID, liked bool) error {
	s.log.Info("[SetTrackLiked] storage started")

	filter := bson.M{"_id": profileId}

	var profile dtoProfile

	err := s.collection.FindOne(ctx, filter, &options.FindOneOptions{}).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return storage.ErrProfileNotFound
		}

		slog.Error("[SetTrackLiked] storage error: " + err.Error())
		return errors.New("[SetTrackLiked] storage error: " + err.Error())
	}

	if liked {
		profile.SavedTracksIds = append(profile.SavedTracksIds, trackId)

	} else {
		addedTracks := 0
		newLikedTracks := make([]primitive.ObjectID, len(profile.SavedTracksIds)-1)
		for i, id := range profile.SavedTracksIds {
			if id != trackId {
				newLikedTracks[addedTracks] = profile.SavedTracksIds[i]
				addedTracks++
			}
		}

		profile.SavedTracksIds = newLikedTracks
	}

	update := bson.M{
		"$set": profile,
	}

	_, err = s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("[SetTrackLiked] storage error: " + err.Error())
		return errors.New("[SetTrackLiked] storage error: " + err.Error())
	}
	return nil
}
