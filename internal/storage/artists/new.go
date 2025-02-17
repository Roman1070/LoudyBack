package artists

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type ArtistsStorage struct {
	collection *mongo.Collection
	log        *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger) *ArtistsStorage {
	storage := &ArtistsStorage{
		collection: db.Collection(collectionName),
		log:        log,
	}

	return storage
}
