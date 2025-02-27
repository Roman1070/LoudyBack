package albums

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *AlbumsStorage) AlbumLight(ctx context.Context, id primitive.ObjectID) (models.AlbumLight, error) {
	s.log.Info("[AlbumLight] storage started")

	filter := bson.M{"_id": id}

	var album dtoAlbum

	err := s.collection.FindOne(ctx, filter).Decode(&album)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.AlbumLight{}, storage.ErrAlbumNotFound
		}

		slog.Error("[Album] storage error: " + err.Error())
		return models.AlbumLight{}, errors.New("[Album] storage error: " + err.Error())
	}

	return models.AlbumLight{
		ID:          album.ID,
		Name:        album.Name,
		Cover:       album.Cover,
		ReleaseDate: album.ReleaseDate,
	}, nil
}
