package tracks

import (
	"context"
	"errors"
	"log/slog"
	tracksv1 "loudy-back/gen/go/tracks"
	models "loudy-back/internal/domain/models/tracks"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Tracks interface {
	CreateTrack(ctx context.Context, name string, albumId primitive.ObjectID, artistsIds []primitive.ObjectID, duration uint16) (primitive.ObjectID, error)
	Track(ctx context.Context, id primitive.ObjectID) (models.Track, error)
	Tracks(ctx context.Context, ids []primitive.ObjectID) ([]models.Track, error)
}

type serverAPI struct {
	tracksv1.UnimplementedTracksServer
	log    *slog.Logger
	tracks Tracks
}

func Register(gRPC *grpc.Server, tracks Tracks, log *slog.Logger) {
	tracksv1.RegisterTracksServer(gRPC, &serverAPI{tracks: tracks, log: log})
}

func (s *serverAPI) CreateTrack(ctx context.Context, req *tracksv1.CreateTrackRequest) (*tracksv1.CreateTrackResponse, error) {
	s.log.Info("[CreateTrack] grpc started")

	albumId, err := primitive.ObjectIDFromHex(req.AlbumId)
	if err != nil {
		s.log.Error("[CreateTrack] grpc error: " + err.Error())
		return nil, errors.New("[CreateTrack] grpc error: " + err.Error())
	}

	artistsIds, err := utils.StringsToIdsArray(req.ArtsitsIds)
	if err != nil {
		s.log.Error("[CreateTrack] grpc error: " + err.Error())
		return nil, errors.New("[CreateTrack] grpc error: " + err.Error())
	}

	resp, err := s.tracks.CreateTrack(ctx, req.Name, albumId, artistsIds, uint16(req.Duration))
	if err != nil {
		s.log.Error("[CreateTrack] grpc error: " + err.Error())
		return nil, errors.New("[CreateTrack] grpc error: " + err.Error())
	}

	return &tracksv1.CreateTrackResponse{
		Id: resp.Hex(),
	}, nil
}

func (s *serverAPI) Track(ctx context.Context, req *tracksv1.TrackRequest) (*tracksv1.TrackResponse, error) {
	s.log.Info("[Track] grpc started")

	trackId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[Track] grpc error: " + err.Error())
		return nil, errors.New("[Track] grpc error: " + err.Error())
	}

	resp, err := s.tracks.Track(ctx, trackId)
	if err != nil {
		s.log.Error("[Track] grpc error: " + err.Error())
		return nil, errors.New("[Track] grpc error: " + err.Error())
	}

	return &tracksv1.TrackResponse{
		Track: resp.ToGRPC(),
	}, nil
}

func (s *serverAPI) Tracks(ctx context.Context, req *tracksv1.TracksRequest) (*tracksv1.TracksResponse, error) {
	s.log.Info("[Tracks] grpc started")

	ids, err := utils.StringsToIdsArray(req.Ids)
	if err != nil {
		s.log.Error("[Tracks] grpc error: " + err.Error())
		return nil, errors.New("[Tracks] grpc error: " + err.Error())
	}

	tracks, err := s.tracks.Tracks(ctx, ids)
	if err != nil {
		s.log.Error("[Tracks] grpc error: " + err.Error())
		return nil, errors.New("[Tracks] grpc error: " + err.Error())
	}

	result := make([]*tracksv1.TrackData, len(tracks))
	for i, track := range tracks {
		result[i] = track.ToGRPC()
	}

	return &tracksv1.TracksResponse{
		Tracks: result,
	}, nil
}
