package artists

import (
	"context"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Artists interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
	Artists(ctx context.Context, ids []primitive.ObjectID) ([]models.Artist, error)
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

func (s *serverAPI) Artist(ctx context.Context, req *artistsv1.ArtistRequest) (*artistsv1.ArtistResponse, error) {
	s.log.Info("[CreateArtist] grpc started")

	artist, err := s.artists.Artist(ctx, req.Name)
	if err != nil {
		s.log.Error("[Artist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	return artist.ToGRPC(), nil
}
func (s *serverAPI) Artists(ctx context.Context, req *artistsv1.ArtistsRequest) (*artistsv1.ArtistsResponse, error) {
	s.log.Info("[CreateArtist] grpc started")

	ids := make([]primitive.ObjectID, len(req.Ids))
	for i, id := range req.Ids {
		newId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			s.log.Error("[Artist] grpc error: " + err.Error())
			return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
		}

		ids[i] = newId
	}

	artists, err := s.artists.Artists(ctx, ids)
	if err != nil {
		s.log.Error("[Artist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	return models.ArtistsToGRPC(artists), nil
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
