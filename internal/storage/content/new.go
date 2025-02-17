package content

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContentStorage struct {
	collection *mongo.Collection
	log        *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger) *ContentStorage {
	storage := &ContentStorage{
		collection: db.Collection(collectionName),
		log:        log,
	}

	return storage
}
