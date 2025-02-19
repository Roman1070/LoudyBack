package artists

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *ArtistsStorage) Artist(ctx context.Context, id primitive.ObjectID) (models.Artist, error) {
	c.log.Info("[Artist] storage started")

	filter := bson.M{"_id": id}

	var result dtoArtist

	err := c.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Artist{}, storage.ErrArtistNotFound
		}

		slog.Error("[Artist] storage error: " + err.Error())
		return models.Artist{}, errors.New("[Artist] storage error: " + err.Error())
	}

	resp, err := c.albumsProvider.AlbumsLight(ctx, result.AlbumsIds)

	if err != nil {
		slog.Error("[Artist] storage error: " + err.Error())
		return models.Artist{}, errors.New("[Artist] storage error: " + err.Error())
	}

	return result.toCommonModel(resp), nil
}
