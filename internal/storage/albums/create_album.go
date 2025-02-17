package albums

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AlbumsStorage) CreateAlbum(ctx context.Context, name, cover string,
	releaseDate string, artistsIds []primitive.ObjectID) (*emptypb.Empty, error) {
	newAlbum := dtoAlbum{
		Name:        name,
		Cover:       cover,
		ReleaseDate: releaseDate,
		ArtistsIds:  artistsIds,
		TracksIds:   []primitive.ObjectID{},
	}

	_, err := s.collection.InsertOne(ctx, newAlbum)
	if err != nil {
		s.log.Error("[CreateArtist] storage error: " + err.Error())
		return nil, errors.New("[CreateArtist] storage error: " + err.Error())
	}

	return nil, nil
}
