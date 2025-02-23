package profiles

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilesStorage struct {
	collection *mongo.Collection
	log        *slog.Logger
}

func NewStorage(db *mongo.Database, collectionName string, log *slog.Logger) *ProfilesStorage {
	storage := &ProfilesStorage{
		collection: db.Collection(collectionName),
		log:        log,
	}

	return storage
}
