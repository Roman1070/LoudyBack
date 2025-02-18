package artists

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ArtistsStorage) DtoArtists(ctx context.Context, ids []primitive.ObjectID) ([]dtoArtist, error) {
	query := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.collection.Find(ctx, query)
	if err != nil {
		s.log.Info("[DtoArtists] cursor error: " + err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, storage.ErrArtistNotFound
		}

		slog.Error("[DtoArtists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[DtoArtists] storage error: "+err.Error())
	}

	var results []dtoArtist
	err = cursor.All(ctx, &results)
	if err != nil {
		slog.Error("[DtoArtists] storage error: " + err.Error())
		return nil, fmt.Errorf("%s", "[DtoArtists] storage error: "+err.Error())
	}

	return results, nil
}
