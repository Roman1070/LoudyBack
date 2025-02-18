package artists

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ArtistsService) AddAlbum(ctx context.Context, artistsIds []primitive.ObjectID, albumId primitive.ObjectID) (*emptypb.Empty, error) {
	s.log.Info("[AddAlbum] service started")
	//TODO: check album existance

	_, err := s.artists.AddAlbum(ctx, artistsIds, albumId)
	if err != nil {
		s.log.Error("[AddAlbum] service error: " + err.Error())
		return nil, errors.New("[AddAlbum] service error: " + err.Error())
	}

	return nil, nil
}
