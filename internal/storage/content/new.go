package content

import "go.mongodb.org/mongo-driver/mongo"

type ContentStorage struct {
	collection *mongo.Collection
}

func NewStorage(db *mongo.Database, collectionName string) *ContentStorage {
	storage := &ContentStorage{
		collection: db.Collection(collectionName),
	}

	return storage
}
