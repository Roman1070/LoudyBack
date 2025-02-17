package content

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ContentService) CreateTrack(ctx context.Context, name, file string, albumId primitive.ObjectID) (*emptypb.Empty, error) {
	return nil, nil
}
