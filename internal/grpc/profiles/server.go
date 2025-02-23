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
	CreateProfile(ctx context.Context, userId uint32, name, avatar, bio string) (id primitive.ObjectID, err error)
	Profile(ctx context.Context, id primitive.ObjectID) (profile models.Profile, err error)
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

func (s *serverAPI) GetProfile(ctx context.Context, req *profilesv1.ProfileRequest) (*profilesv1.ProfileResponse, error) {
	s.log.Info("[GetProfile] grpc started")

	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		s.log.Error("[GetProfile] grpc error: " + err.Error())
		return nil, errors.New("[GetProfile] grpc error: " + err.Error())
	}

	profile, err := s.profiles.Profile(ctx, id)
	if err != nil {
		s.log.Error("[GetProfile] grpc error: " + err.Error())
		return nil, errors.New("[GetProfile] grpc error: " + err.Error())
	}

	return &profilesv1.ProfileResponse{
		Profile: profile.ToGRPC(),
	}, nil
}
