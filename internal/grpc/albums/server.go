package albums

import (
	"context"
	"errors"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/albums"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Albums interface {
	Album(ctx context.Context, id string) (models.Album, error)
	CreateAlbum(ctx context.Context, name, cover string, artists_ids []primitive.ObjectID) (*emptypb.Empty, error)
}
type serverAPI struct {
	albums Albums
	log    *slog.Logger
	albumsv1.UnimplementedAlbumsServer
}

func Register(gRPC *grpc.Server, albums Albums, log *slog.Logger) {
	albumsv1.RegisterAlbumsServer(gRPC, &serverAPI{albums: albums, log: log})
}

func (s *serverAPI) Album(ctx context.Context, req *albumsv1.AlbumRequest) (*albumsv1.AlbumResponse, error) {
	s.log.Info("[Album] grpc started")

	album, err := s.albums.Album(ctx, req.Id)
	if err != nil {
		s.log.Error("[Album] grpc error: " + err.Error())
		return nil, errors.New("[Album] grpc error: " + err.Error())
	}

	return album.ToGRPC(), nil
}

func (s *serverAPI) CreateAlbum(ctx context.Context, req *albumsv1.CreateAlbumRequest) (*emptypb.Empty, error) {
	s.log.Info("[CreateAlbum] grpc started")

	ids := make([]primitive.ObjectID, len(req.ArtistsIds))
	for i, id := range req.ArtistsIds {
		newId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			s.log.Error("[CreateAlbum] grpc error: " + err.Error())
			return nil, errors.New("[CreateAlbum] grpc error: " + err.Error())
		}

		ids[i] = newId
	}
	_, err := s.albums.CreateAlbum(ctx, req.Name, req.Cover, ids)
	if err != nil {
		s.log.Error("[CreateAlbum] grpc error: " + err.Error())
		return nil, errors.New("[CreateAlbum] grpc error: " + err.Error())
	}

	return nil, nil
}
