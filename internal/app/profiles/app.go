package app

import (
	"fmt"
	"log/slog"
	mongo_db "loudy-back/configs/mongo"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/profiles"
	"loudy-back/internal/services/profiles"
	repositoryProfiles "loudy-back/internal/storage/profiles"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	albumsProvider profiles.AlbumsProvider,
	artistsProvider profiles.ArtistsProvider,
	tracksProvider profiles.TracksProvider,
) (*App, error) {

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := repositoryProfiles.NewStorage(mongoDb, "profiles", log)

	profilesService := profiles.New(log, repo, artistsProvider, albumsProvider, tracksProvider)
	grpcApp := grpcApp.New(log, profilesService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
