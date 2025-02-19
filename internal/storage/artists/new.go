package artists

import (
	"context"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumsProvider interface {
	AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error)
}

type ArtistsStorage struct {
	collection     *mongo.Collection
	albumsProvider AlbumsProvider
	log            *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger, albumsProvider AlbumsProvider) *ArtistsStorage {
	storage := &ArtistsStorage{
		collection:     db.Collection(collectionName),
		albumsProvider: albumsProvider,
		log:            log,
	}

	return storage
}
