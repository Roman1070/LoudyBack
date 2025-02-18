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
	//TODO: попробовать за один UpdateMany сделать
	for _, artist := range artists {
		artist.AlbumsIds = append(artist.AlbumsIds, albumId)

		filter := bson.M{"_id": artist.ID}
		update := bson.M{
			"$set": artist,
		}

		_, err = s.collection.UpdateOne(ctx, filter, update)
		if err != nil {
			s.log.Error("[Artists] storage error: " + err.Error())
			return nil, errors.New("[Artists] storage error: " + err.Error())
		}
	}

	return nil, nil
}
