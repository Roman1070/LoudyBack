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
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
	AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error)
	CreateAlbum(ctx context.Context, name, cover string, releaseDate string,
		artists_ids []primitive.ObjectID) (primitive.ObjectID, error)
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

	artists := make([]*albumsv1.ArtistLight, len(album.Artists))
	for i, artist := range album.Artists {
		artists[i] = &albumsv1.ArtistLight{
			Id:   artist.ID.Hex(),
			Name: artist.Name,
		}
	}

	tracks := make([]*albumsv1.TrackLight, len(album.Tracks))
	for i, track := range album.Tracks {
		tracks[i] = &albumsv1.TrackLight{
			Id:       track.ID.Hex(),
			Name:     track.Name,
			Duration: uint32(track.Duration),
		}
	}

	return &albumsv1.AlbumResponse{
		Id:          album.ID.Hex(),
		Name:        album.Name,
		Cover:       album.Cover,
		ReleaseDate: album.ReleaseDate,
		Artists:     artists,
		Tracks:      tracks,
	}, nil
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

func (s *serverAPI) AlbumsLight(ctx context.Context, req *albumsv1.AlbumsLightRequest) (*albumsv1.AlbumsLightResponse, error) {
	s.log.Info("[AlbumsLight] grpc started")

	ids, err := utils.StringsToIdsArray(req.Ids)
	if err != nil {
		s.log.Error("[AlbumsLight] grpc error: " + err.Error())
		return nil, errors.New("[AlbumsLight] grpc error: " + err.Error())
	}

	albums, err := s.albums.AlbumsLight(ctx, ids)
	if err != nil {
		s.log.Error("[AlbumsLight] grpc error: " + err.Error())
		return nil, errors.New("[AlbumsLight] grpc error: " + err.Error())
	}

	return models.AlbumsLightToGRPC(albums), nil
}
