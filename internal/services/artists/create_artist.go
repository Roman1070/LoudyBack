package artists

import (
	"context"
	"fmt"
	"loudy-back/internal/storage"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ArtistsService) CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error) {
	s.log.Info("[CreateArtist] service started")

	_, err := s.artistsProvider.ArtistLightByName(ctx, name)
	if err == nil {
		s.log.Error("[CreateArtist] service error: Artist already exists")
		return nil, storage.ErrArtistAlreadyExists
	}

	id, err := s.artists.CreateArtist(ctx, name, cover, bio)
	if err != nil {
		s.log.Error("[CreateArtist] service error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateArtist] service error: "+err.Error())
	}

	s.log.Info("[CreateArtist] finished successfully, id=" + fmt.Sprint(id))
	return nil, nil
}
