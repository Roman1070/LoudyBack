package app

import (
	"fmt"
	"log/slog"
	mongo_db "loudy-back/configs/mongo"
	albumsv1 "loudy-back/gen/go/albums"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/artists"
	"loudy-back/internal/services/artists"
	repositoryArtists "loudy-back/internal/storage/artists"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	albumsClient albumsv1.AlbumsClient,
	albumsProvider repositoryArtists.AlbumsProvider,
) (*App, error) {

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := repositoryArtists.NewStorage(mongoDb, "artists", log, albumsProvider)

	artistsService := artists.New(log, repo, albumsClient, repo)

	grpcApp := grpcApp.New(log, artistsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
