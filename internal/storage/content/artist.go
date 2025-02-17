package content

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/content"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *ContentStorage) Artist(ctx context.Context, name string) (models.Artist, error) {
	c.log.Info("[Artist] storage started")

	filter := bson.M{"name": name}

	var result dtoArtist

	err := c.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Artist{}, storage.ErrArtistNotFound
		}

		slog.Error("[Artist] storage error: " + err.Error())
		return models.Artist{}, fmt.Errorf("%s", "[Artist] storage error: "+err.Error())
	}

	return result.toCommonModel(), nil
}
