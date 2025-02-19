package artists

import (
	"context"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Artists interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
	AddAlbum(ctx context.Context, artistId []primitive.ObjectID, albumId primitive.ObjectID) (*emptypb.Empty, error)
}

type serverAPI struct {
	artistsv1.UnimplementedArtistsServer
	log     *slog.Logger
	artists Artists
}

func Register(gRPC *grpc.Server, artists Artists, log *slog.Logger) {
	artistsv1.RegisterArtistsServer(gRPC, &serverAPI{artists: artists, log: log})
}

func (s *serverAPI) AddAlbum(ctx context.Context, req *artistsv1.AddAlbumRequest) (*emptypb.Empty, error) {
	s.log.Info("[AddAlbum] grpc started")

	artistIds, err := utils.StringsToIdsArray(req.ArtistsIds)
	if err != nil {
		s.log.Error("[AddAlbum] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[AddAlbum] grpc error: "+err.Error())
	}

	albumId, err := primitive.ObjectIDFromHex(req.AlbumId)
	if err != nil {
		s.log.Error("[AddAlbum] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[AddAlbum] grpc error: "+err.Error())
	}

	_, err = s.artists.AddAlbum(ctx, artistIds, albumId)
	if err != nil {
		s.log.Error("[AddAlbum] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[AddAlbum] grpc error: "+err.Error())
	}

	return nil, nil
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
