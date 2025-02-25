package tracks

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type TracksStorage struct {
	collection *mongo.Collection
	log        *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger) *TracksStorage {
	storage := &TracksStorage{
		collection: db.Collection(collectionName),
		log:        log,
	}

	return storage
}
