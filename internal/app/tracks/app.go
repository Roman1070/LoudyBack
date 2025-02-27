package app

import (
	"fmt"
	"log/slog"
	mongo_db "loudy-back/configs/mongo"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/tracks"
	"loudy-back/internal/services/tracks"
	repositorytracks "loudy-back/internal/storage/tracks"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	albumsProvider tracks.AlbumsProvider,
	artistsProvider tracks.ArtistsProvider,
) (*App, error) {

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := repositorytracks.NewStorage(mongoDb, "tracks", log)

	tracksService := tracks.New(repo, artistsProvider, albumsProvider, log)
	grpcApp := grpcApp.New(log, tracksService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
