package artists

import (
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"

	"go.mongodb.org/mongo-driver/mongo"
)

type ArtistsStorage struct {
	collection   *mongo.Collection
	albumsClient albumsv1.AlbumsClient
	log          *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger, albumsClient albumsv1.AlbumsClient) *ArtistsStorage {
	storage := &ArtistsStorage{
		collection:   db.Collection(collectionName),
		albumsClient: albumsClient,
		log:          log,
	}

	return storage
}
