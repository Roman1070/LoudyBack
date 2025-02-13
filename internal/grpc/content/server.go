package content

import (
	"context"
	contentv1 "loudy-back/gen/go/content"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Content interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (uint32, error)
	CreateTrack(ctx context.Context, name, file string, albumId uint32) (*emptypb.Empty, error)
	CreateAlbum(ctx context.Context, name, cover string, tracksIds []uint32, releaseDate time.Time) (uint32, error)
}

type serverAPI struct {
	contentv1.UnimplementedContentServer
	content Content
}

func Register(gRPC *grpc.Server, content Content) {
	contentv1.RegisterContentServer(gRPC, &serverAPI{content: content})
}

func (s *serverAPI) CreateAlbum(ctx context.Context, req *contentv1.CreateAlbumRequest) (*contentv1.CreateAlbumResponse, error) {
	return nil, nil
}

func (s *serverAPI) CreateArtist(ctx context.Context, req *contentv1.CreateArtistRequest) (*contentv1.CreateArtistResponse, error) {
	return nil, nil
}
func (s *serverAPI) CreateTrack(ctx context.Context, req *contentv1.CreateTrackRequest) (*emptypb.Empty, error) {
	return nil, nil
}
