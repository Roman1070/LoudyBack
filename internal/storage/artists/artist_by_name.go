package artists

import (
	"context"
	"errors"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	albumModels "loudy-back/internal/domain/models/albums"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ArtistsStorage) ArtistByName(ctx context.Context, name string) (models.Artist, error) {
	s.log.Info("[ArtistByName] storage started")

	filter := bson.M{"name": name}

	var result dtoArtist

	err := s.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Artist{}, storage.ErrArtistNotFound
		}

		slog.Error("[ArtistByName] storage error: " + err.Error())
		return models.Artist{}, errors.New("[ArtistByName] storage error: " + err.Error())
	}

	resp, err := s.albumsClient.AlbumsLight(ctx, &albumsv1.AlbumsLightRequest{
		Ids: utils.IdsToStringArray(result.AlbumsIds),
	})

	if err != nil {
		slog.Error("[Artist] storage error: " + err.Error())
		return models.Artist{}, errors.New("[ArtistByName] storage error: " + err.Error())
	}

	albums, err := albumModels.GRPCResponseToAlbumsLight(resp.Albums)
	if err != nil {
		slog.Error("[ArtistByName] storage error: " + err.Error())
		return models.Artist{}, errors.New("[ArtistByName] storage error: " + err.Error())
	}

	return result.toCommonModel(albums), nil
}
