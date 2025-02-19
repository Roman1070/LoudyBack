package albums

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Albums interface {
	CreateAlbum(ctx context.Context, name, cover string, releaseDate string,
		artists_ids []primitive.ObjectID) (primitive.ObjectID, error)
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
}

type serverAPI struct {
	albums Albums
	log    *slog.Logger
	albumsv1.UnimplementedAlbumsServer
}

func Register(gRPC *grpc.Server, albums Albums, log *slog.Logger) {
	albumsv1.RegisterAlbumsServer(gRPC, &serverAPI{albums: albums, log: log})
}

func (s *serverAPI) CreateAlbum(ctx context.Context, req *albumsv1.CreateAlbumRequest) (*albumsv1.CreateAlbumResponse, error) {
	s.log.Info("[CreateAlbum] grpc started")

	ids, err := utils.StringsToIdsArray(req.ArtistsIds)
	if err != nil {
		s.log.Error("[CreateAlbum] grpc error: " + err.Error())
		return nil, errors.New("[CreateAlbum] grpc error: " + err.Error())
	}

	resp, err := s.albums.CreateAlbum(ctx, req.Name, req.Cover, req.ReleaseDate, ids)
	if err != nil {
		s.log.Error("[CreateAlbum] grpc error: " + err.Error())
		return nil, errors.New("[CreateAlbum] grpc error: " + err.Error())
	}

	return &albumsv1.CreateAlbumResponse{
		Id: resp.Hex(),
	}, nil
}

func (s *serverAPI) Album(ctx context.Context, req *albumsv1.AlbumRequest) (*albumsv1.AlbumResponse, error) {
	s.log.Info("[Album] grpc started")

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[Album] grpc error: " + err.Error())
		return nil, errors.New("[Album] grpc error: " + err.Error())
	}

	album, err := s.albums.Album(ctx, id)
	if err != nil {
		s.log.Error("[Album] grpc error: " + err.Error())
		return nil, errors.New("[Album] grpc error: " + err.Error())
	}

	s.log.Info("[Album] grpc recieved album: " + fmt.Sprint(album))
	return &albumsv1.AlbumResponse{
		Album: album.ToGRPC(),
	}, nil
}
