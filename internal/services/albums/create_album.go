package albums

import (
	"context"
	"fmt"
	"loudy-back/internal/storage"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *AlbumsService) CreateAlbum(ctx context.Context, name, cover string, artistsIds []primitive.ObjectID) (*emptypb.Empty, error) {
	s.log.Info("[CreateAlbum] service started")

	_, err := s.artists.Artist(ctx, name)
	if err == nil {
		s.log.Error("[CreateAlbum] service error: Artist already exists")
		return nil, storage.ErrArtistAlreadyExists
	}

	_, err = s.albums.CreateAlbum(ctx, name, cover, artistsIds)
	if err != nil {
		s.log.Error("[CreateAlbum] service error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateAlbum] service error: "+err.Error())
	}

	return nil, nil
}
