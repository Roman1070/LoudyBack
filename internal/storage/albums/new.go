package albums

import (
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"

	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumsStorage struct {
	collection    *mongo.Collection
	artistsClient artistsv1.ArtistsClient
	log           *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, artistsClient artistsv1.ArtistsClient, log *slog.Logger) *AlbumsStorage {
	storage := &AlbumsStorage{
		collection:    db.Collection(collectionName),
		artistsClient: artistsClient,
		log:           log,
	}

	return storage
}
