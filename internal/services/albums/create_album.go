package albums

import (
	"context"
	"fmt"
	artistsv1 "loudy-back/gen/go/artists"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) CreateAlbum(ctx context.Context, name, cover string,
	releaseDate string, artistsIds []primitive.ObjectID) (primitive.ObjectID, error) {

	s.log.Info("[CreateAlbum] service started")
	//TODO: check album existance

	resp, err := s.albums.CreateAlbum(ctx, name, cover, releaseDate, artistsIds)
	if err != nil {
		s.log.Error("[CreateAlbum] service error: " + err.Error())
		return [12]byte{}, fmt.Errorf("%s", "[CreateAlbum] service error: "+err.Error())
	}

	_, err = s.artists.AddAlbum(ctx, &artistsv1.AddAlbumRequest{
		ArtistsIds: utils.IdsToStringArray(artistsIds),
		AlbumId:    resp.Hex(),
	})
	if err != nil {
		s.log.Error("[CreateAlbum] service error: " + err.Error())
		return [12]byte{}, fmt.Errorf("%s", "[CreateAlbum] service error: "+err.Error())
	}

	return resp, nil
}
