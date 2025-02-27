package app

import (
	"fmt"
	"log/slog"
	mongo_db "loudy-back/configs/mongo"
	artistsv1 "loudy-back/gen/go/artists"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/albums"
	"loudy-back/internal/services/albums"
	repositoryAlbums "loudy-back/internal/storage/albums"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	artistsClient artistsv1.ArtistsClient,
	artistsProvider albums.ArtistsProvider,
	tracksProvider albums.TracksProvider,
) (*App, error) {

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := repositoryAlbums.NewStorage(mongoDb, "albums", artistsClient, log)

	albumsService := albums.New(log, artistsClient, repo, artistsProvider, tracksProvider)

	grpcApp := grpcApp.New(log, albumsService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
