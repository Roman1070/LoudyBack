package content

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ContentStorage) CreateTrack(ctx context.Context, name, file string, albumId uint32) (*emptypb.Empty, error) {
	return nil, nil
}
