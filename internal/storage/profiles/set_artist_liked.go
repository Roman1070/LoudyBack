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

func (s *ProfilesStorage) SetArtistLiked(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID, liked bool) error {
	s.log.Info("[SetArtistLiked] storage started")

	filter := bson.M{"_id": profileId}

	var profile dtoProfile

	err := s.collection.FindOne(ctx, filter, &options.FindOneOptions{}).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return storage.ErrProfileNotFound
		}

		slog.Error("[SetArtistLiked] storage error: " + err.Error())
		return errors.New("[SetArtistLiked] storage error: " + err.Error())
	}

	if liked {
		profile.SavedArtistsIds = append(profile.SavedArtistsIds, artistId)

	} else {
		addedArtists := 0
		newLikedArtists := make([]primitive.ObjectID, len(profile.SavedArtistsIds)-1)
		for i, id := range profile.SavedArtistsIds {
			if id != artistId {
				newLikedArtists[addedArtists] = profile.SavedArtistsIds[i]
				addedArtists++
			}
		}

		profile.SavedArtistsIds = newLikedArtists
	}

	update := bson.M{
		"$set": profile,
	}

	_, err = s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("[SetArtistLiked] storage error: " + err.Error())
		return errors.New("[SetArtistLiked] storage error: " + err.Error())
	}
	return nil
}
