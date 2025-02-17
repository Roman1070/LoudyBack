package content

import (
	"context"
	"fmt"
	"log/slog"
	contentv1 "loudy-back/gen/go/content"
	models "loudy-back/internal/domain/models/content"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Content interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (*emptypb.Empty, error)
	CreateTrack(ctx context.Context, name, file string, albumId primitive.ObjectID) (*emptypb.Empty, error)
	CreateAlbum(ctx context.Context, name, cover string, tracks []models.TrackLight, releaseDate time.Time) (*emptypb.Empty, error)
}

type serverAPI struct {
	contentv1.UnimplementedContentServer
	log     *slog.Logger
	content Content
}

func Register(gRPC *grpc.Server, content Content, log *slog.Logger) {
	contentv1.RegisterContentServer(gRPC, &serverAPI{content: content, log: log})
}

func (s *serverAPI) CreateAlbum(ctx context.Context, req *contentv1.CreateAlbumRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *serverAPI) CreateArtist(ctx context.Context, req *contentv1.CreateArtistRequest) (*emptypb.Empty, error) {
	s.log.Info("[CreateArtist] grpc started")

	_, err := s.content.CreateArtist(ctx, req.Name, req.Cover, req.Bio)
	if err != nil {
		s.log.Error("[CreateArtist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[CreateArtist] grpc error: "+err.Error())
	}

	return nil, nil
}
func (s *serverAPI) CreateTrack(ctx context.Context, req *contentv1.CreateTrackRequest) (*emptypb.Empty, error) {
	return nil, nil
}
