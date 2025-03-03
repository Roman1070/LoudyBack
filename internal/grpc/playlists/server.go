package playlists

import (
	"context"
	"errors"
	"log/slog"
	playlistsv1 "loudy-back/gen/go/playlists"
	models "loudy-back/internal/domain/models/playlists"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Playlists interface {
	Playlist(ctx context.Context, id primitive.ObjectID) (models.Playlist, error)
}

type serverAPI struct {
	playlistsv1.UnimplementedPlaylistsServer
	playlists Playlists
	log       *slog.Logger
}

func Register(gRPC *grpc.Server, playlists Playlists) {
	playlistsv1.RegisterPlaylistsServer(gRPC, &serverAPI{playlists: playlists})
}

func (s *serverAPI) Playlist(ctx context.Context, req *playlistsv1.PlaylistRequest) (*playlistsv1.PlaylistResponse, error) {
	s.log.Info("[Playlist] grpc started")

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[Playlist] grpc error: " + err.Error())
		return nil, errors.New("[Playlist] grpc error: " + err.Error())
	}

	resp, err := s.playlists.Playlist(ctx, id)
	if err != nil {
		s.log.Error("[Playlist] grpc error: " + err.Error())
		return nil, errors.New("[Playlist] grpc error: " + err.Error())
	}

	return &playlistsv1.PlaylistResponse{
		Data: resp.ToGRPC(),
	}, nil
}
