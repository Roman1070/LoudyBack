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

func (c *ArtistsStorage) Artists(ctx context.Context, ids []primitive.ObjectID) ([]models.Artist, error) {
	c.log.Info("[Artists] storage started")

	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := c.collection.Find(ctx, query)
	if err != nil {
		c.log.Info("[Artists] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrArtistNotFound
		}

		slog.Error("[Artists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artists] storage error: "+err.Error())
	}
	c.log.Info("[Artists] cursor recieved")

	var results []dtoArtist
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[Artists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artists] storage error: "+err.Error())
	}
	c.log.Info("[Artists] results written, results= " + fmt.Sprint(results))

	return toCommonModels(results), nil
}
