package artists

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ArtistsStorage) AddAlbum(ctx context.Context, artistsIds []primitive.ObjectID, albumId primitive.ObjectID) (*emptypb.Empty, error) {
	s.log.Info("[AddAlbum] storge started")

	artists, err := s.DtoArtists(ctx, artistsIds)
	if err != nil {
		s.log.Error("[Artists] storage error: " + err.Error())
		return nil, errors.New("[Artists] storage error: " + err.Error())
	}

	for _, artist := range artists {
		artist.AlbumsIds = append(artist.AlbumsIds, albumId)
	}

	filter := bson.M{"_id": bson.M{"$in": artistsIds}}
	_, err = s.collection.UpdateMany(ctx, filter, artists)
	if err != nil {
		s.log.Error("[Artists] storage error: " + err.Error())
		return nil, errors.New("[Artists] storage error: " + err.Error())
	}

	return nil, nil
}
