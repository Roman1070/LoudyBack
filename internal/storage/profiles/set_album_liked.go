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

func (s *ProfilesStorage) SetAlbumLiked(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID, liked bool) error {
	s.log.Info("[SetAlbumLiked] storage started")

	filter := bson.M{"_id": profileId}

	var profile dtoProfile

	err := s.collection.FindOne(ctx, filter, &options.FindOneOptions{}).Decode(&profile)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return storage.ErrProfileNotFound
		}

		slog.Error("[SetAlbumLiked] storage error: " + err.Error())
		return errors.New("[SetAlbumLiked] storage error: " + err.Error())
	}

	if liked {
		profile.SavedAlbumsIds = append(profile.SavedAlbumsIds, albumId)

	} else {
		addedAlbums := 0
		newLikedAlbums := make([]primitive.ObjectID, len(profile.SavedAlbumsIds)-1)
		for i, id := range profile.SavedAlbumsIds {
			if id != albumId {
				newLikedAlbums[addedAlbums] = profile.SavedAlbumsIds[i]
				addedAlbums++
			}
		}

		profile.SavedAlbumsIds = newLikedAlbums
	}

	update := bson.M{
		"$set": profile,
	}

	_, err = s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("[SetAlbumLiked] storage error: " + err.Error())
		return errors.New("[SetAlbumLiked] storage error: " + err.Error())
	}
	return nil
}
