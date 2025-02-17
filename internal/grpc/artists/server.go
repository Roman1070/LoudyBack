package artists

import (
	"context"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Artists interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
}

type serverAPI struct {
	artistsv1.UnimplementedArtistsServer
	log     *slog.Logger
	artists Artists
}

func Register(gRPC *grpc.Server, artists Artists, log *slog.Logger) {
	artistsv1.RegisterArtistsServer(gRPC, &serverAPI{artists: artists, log: log})
}

func (s *serverAPI) CreateArtist(ctx context.Context, req *artistsv1.CreateArtistRequest) (*emptypb.Empty, error) {
	s.log.Info("[CreateArtist] grpc started")

	_, err := s.artists.CreateArtist(ctx, req.Name, req.Cover, req.Bio)
	if err != nil {
		s.log.Error("[CreateArtist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateArtist] grpc error: "+err.Error())
	}

	return nil, nil
}
