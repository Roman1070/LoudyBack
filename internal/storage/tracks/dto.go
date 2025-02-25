package tracks

import "go.mongodb.org/mongo-driver/bson/primitive"

type dtoTrack struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Name       string               `bson:"name"`
	AlbumId    primitive.ObjectID   `bson:"album_id"`
	ArtistsIds []primitive.ObjectID `bson:"artists_ids"`
	Duration   uint16               `bson:"duration"`
}
