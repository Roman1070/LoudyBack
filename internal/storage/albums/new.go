package albums

import (
	"context"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArtistsProvider interface {
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error)
}

type AlbumsStorage struct {
	collection      *mongo.Collection
	artistsClient   artistsv1.ArtistsClient
	artistsProvider ArtistsProvider
	log             *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, artistsClient artistsv1.ArtistsClient, artistsProvider ArtistsProvider, log *slog.Logger) *AlbumsStorage {
	storage := &AlbumsStorage{
		collection:      db.Collection(collectionName),
		artistsClient:   artistsClient,
		log:             log,
		artistsProvider: artistsProvider,
	}

	return storage
}
