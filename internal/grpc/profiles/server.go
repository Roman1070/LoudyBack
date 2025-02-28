package profiles

import (
	"context"
	"errors"
	"log/slog"
	profilesv1 "loudy-back/gen/go/profiles"
	models "loudy-back/internal/domain/models/profiles"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Profiles interface {
	CreateProfile(ctx context.Context, userId uint32, name, avatar, bio string) (primitive.ObjectID, error)
	Profile(ctx context.Context, id primitive.ObjectID) (models.Profile, error)
	ToggleLikeTrack(ctx context.Context, profileId primitive.ObjectID, trackId primitive.ObjectID) (bool, error)
	ToggleLikeAlbum(ctx context.Context, profileId primitive.ObjectID, albumId primitive.ObjectID) (bool, error)
	ToggleLikeArtist(ctx context.Context, profileId primitive.ObjectID, artistId primitive.ObjectID) (bool, error)
}

type serverAPI struct {
	profilesv1.UnimplementedProfilesServer
	log      *slog.Logger
	profiles Profiles
}

func Register(gRPC *grpc.Server, profiles Profiles, log *slog.Logger) {
	profilesv1.RegisterProfilesServer(gRPC, &serverAPI{profiles: profiles, log: log})
}

func (s *serverAPI) CreateProfile(ctx context.Context, req *profilesv1.CreateProfileRequest) (*profilesv1.CreateProfileResponse, error) {
	s.log.Info("[CreateProfile] grpc started")

	id, err := s.profiles.CreateProfile(ctx, req.UserId, req.Username, req.Avatar, req.Bio)
	if err != nil {
		s.log.Error("[CreateProfile] grpc error: " + err.Error())
		return nil, errors.New("[CreateProfile] grpc error: " + err.Error())
	}

	return &profilesv1.CreateProfileResponse{
		Id: id.Hex(),
	}, nil
}

func (s *serverAPI) Profile(ctx context.Context, req *profilesv1.ProfileRequest) (*profilesv1.ProfileResponse, error) {
	s.log.Info("[Profile] grpc started")

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[Profile] grpc error: " + err.Error())
		return nil, errors.New("[Profile] grpc error: " + err.Error())
	}

	profile, err := s.profiles.Profile(ctx, id)
	if err != nil {
		s.log.Error("[Profile] grpc error: " + err.Error())
		return nil, errors.New("[Profile] grpc error: " + err.Error())
	}

	return &profilesv1.ProfileResponse{
		Profile: profile.ToGRPC(),
	}, nil
}

func (s *serverAPI) ToggleLikeTrack(ctx context.Context, req *profilesv1.ToggleLikeTrackRequest) (*profilesv1.ToggleLikeTrackResponse, error) {
	s.log.Info("[ToggleLikeTrack] grpc started")

	profileId, err := primitive.ObjectIDFromHex(req.ProfileId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	trackId, err := primitive.ObjectIDFromHex(req.TrackId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	liked, err := s.profiles.ToggleLikeTrack(ctx, profileId, trackId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	return &profilesv1.ToggleLikeTrackResponse{
		NowLiked: liked,
	}, nil
}

func (s *serverAPI) ToggleLikeAlbum(ctx context.Context, req *profilesv1.ToggleLikeAlbumRequest) (*profilesv1.ToggleLikeAlbumResponse, error) {
	s.log.Info("[ToggleLikeAlbum] grpc started")

	profileId, err := primitive.ObjectIDFromHex(req.ProfileId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	albumId, err := primitive.ObjectIDFromHex(req.AlbumId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	liked, err := s.profiles.ToggleLikeAlbum(ctx, profileId, albumId)
	if err != nil {
		s.log.Error("[ToggleLikeAlbum] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeAlbum] grpc error: " + err.Error())
	}

	return &profilesv1.ToggleLikeAlbumResponse{
		NowLiked: liked,
	}, nil
}

func (s *serverAPI) ToggleLikeArtist(ctx context.Context, req *profilesv1.ToggleLikeArtistRequest) (*profilesv1.ToggleLikeArtistResponse, error) {
	s.log.Info("[ToggleLikeArtist] grpc started")

	profileId, err := primitive.ObjectIDFromHex(req.ProfileId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	artistId, err := primitive.ObjectIDFromHex(req.ArtistId)
	if err != nil {
		s.log.Error("[ToggleLikeTrack] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeTrack] grpc error: " + err.Error())
	}

	liked, err := s.profiles.ToggleLikeArtist(ctx, profileId, artistId)
	if err != nil {
		s.log.Error("[ToggleLikeArtist] grpc error: " + err.Error())
		return nil, errors.New("[ToggleLikeArtist] grpc error: " + err.Error())
	}

	return &profilesv1.ToggleLikeArtistResponse{
		NowLiked: liked,
	}, nil
}
