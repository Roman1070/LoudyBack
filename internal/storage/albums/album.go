package albums

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"
	trackModels "loudy-back/internal/domain/models/tracks"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *AlbumsStorage) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	c.log.Info("[Album] storage started")

	filter := bson.M{"_id": id}

	var result dtoAlbum

	err := c.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Album{}, storage.ErrAlbumNotFound
		}

		slog.Error("[Album] storage error: " + err.Error())
		return models.Album{}, errors.New("[Album] storage error: " + err.Error())
	}

	artists, err := c.artistsProvider.ArtistsLight(ctx, result.ArtistsIds)
	if err != nil {
		slog.Error("[Album] storage error: " + err.Error())
		return models.Album{}, errors.New("[Album] storage error: " + err.Error())
	}

	c.log.Info("[Album] storage finished, result: " + fmt.Sprint(result))
	return result.toCommonModel(artists, []trackModels.TrackLight{}), nil
}
