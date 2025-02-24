package artists

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *ArtistsStorage) ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error) {
	c.log.Info("[ArtistsLight] storage started")
	if len(ids) == 0 {
		return []models.ArtistLight{}, nil
	}

	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := c.collection.Find(ctx, query)
	if err != nil {
		c.log.Info("[ArtistsLight] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrArtistNotFound
		}

		slog.Error("[ArtistsLight] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[ArtistsLight] storage error: "+err.Error())
	}
	c.log.Info("[Artists] cursor recieved")

	var results []dtoArtistLight
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[Artists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artists] storage error: "+err.Error())
	}

	return toCommonModels(results), nil
}
