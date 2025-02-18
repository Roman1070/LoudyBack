package albums

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AlbumsService) CreateAlbum(ctx context.Context, name, cover string,
	releaseDate string, artistsIds []primitive.ObjectID) (*emptypb.Empty, error) {

	s.log.Info("[CreateAlbum] service started")
	//TODO: check album existance

	_, err := s.albums.CreateAlbum(ctx, name, cover, releaseDate, artistsIds)
	if err != nil {
		s.log.Error("[CreateAlbum] service error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateAlbum] service error: "+err.Error())
	}

	return nil, nil
}
