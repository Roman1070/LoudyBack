package albums

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsStorage) CreateAlbum(ctx context.Context, name, cover string,
	releaseDate string, artistsIds []primitive.ObjectID) (primitive.ObjectID, error) {
	newAlbum := dtoAlbum{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Cover:       cover,
		ReleaseDate: releaseDate,
		ArtistsIds:  artistsIds,
		TracksIds:   []primitive.ObjectID{},
	}

	_, err := s.collection.InsertOne(ctx, newAlbum)
	if err != nil {
		s.log.Error("[CreateArtist] storage error: " + err.Error())
		return [12]byte{}, errors.New("[CreateArtist] storage error: " + err.Error())
	}

	return newAlbum.ID, nil
}
