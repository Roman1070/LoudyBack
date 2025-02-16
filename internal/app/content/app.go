package app

import (
	"fmt"
	"log/slog"
	mongo_db "loudy-back/configs/mongo"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/content"
	"loudy-back/internal/services/content"
	repositoryContent "loudy-back/internal/storage/content"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) (*App, error) {

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return nil, fmt.Errorf("[ ERROR ] не инициализируется монго %v", err)
	}

	repo := repositoryContent.NewStorage(mongoDb, "artists")

	contentService := content.New(log, repo)

	grpcApp := grpcApp.New(log, contentService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
