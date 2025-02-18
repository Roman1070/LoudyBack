package artists

import (
	"context"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Artists interface {
	Artist(ctx context.Context, id primitive.ObjectID) (models.Artist, error)
	ArtistByName(ctx context.Context, name string) (models.Artist, error)
	ArtistsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.ArtistLight, error)
	ArtistLightByName(ctx context.Context, name string) (models.ArtistLight, error)
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
func (s *serverAPI) Artist(ctx context.Context, req *artistsv1.ArtistRequest) (*artistsv1.ArtistResponse, error) {
	s.log.Info("[Artist] grpc started")

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[Artist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	artist, err := s.artists.Artist(ctx, id)
	if err != nil {
		s.log.Error("[Artist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	return artist.ToGRPC(), nil
}
func (s *serverAPI) ArtistByName(ctx context.Context, req *artistsv1.ArtistByNameRequest) (*artistsv1.ArtistByNameResponse, error) {
	s.log.Info("[ArtistByName] grpc started")

	artist, err := s.artists.ArtistByName(ctx, req.Name)
	if err != nil {
		s.log.Error("[Artist] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	return &artistsv1.ArtistByNameResponse{
		Artist: artist.ToGRPC().Artist,
	}, nil
}
func (s *serverAPI) ArtistLightByName(ctx context.Context, req *artistsv1.ArtistLightByNameRequest) (*artistsv1.ArtistLightByNameResponse, error) {
	s.log.Info("[ArtistLightByName] grpc started")

	artist, err := s.artists.ArtistLightByName(ctx, req.Name)
	if err != nil {
		s.log.Error("[ArtistLightByName] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[ArtistLightByName] grpc error: "+err.Error())
	}

	return &artistsv1.ArtistLightByNameResponse{
		Artist: &artistsv1.ArtistLight{
			Id:   artist.ID.Hex(),
			Name: artist.Name,
		},
	}, nil
}
func (s *serverAPI) ArtistsLight(ctx context.Context, req *artistsv1.ArtistsLightRequest) (*artistsv1.ArtistsLightResponse, error) {
	s.log.Info("[ArtistsLight] grpc started")

	ids, err := utils.StringsToIdsArray(req.Ids)
	if err != nil {
		s.log.Error("[ArtistsLight] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	artists, err := s.artists.ArtistsLight(ctx, ids)
	if err != nil {
		s.log.Error("[ArtistsLight] grpc error: " + err.Error())
		return nil, fmt.Errorf("%s", "[Artist] grpc error: "+err.Error())
	}

	s.log.Info("[ArtistsLight] grpc artists recieved, artists= " + fmt.Sprint(artists))
	return models.ArtistsLightToGRPC(artists), nil
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
