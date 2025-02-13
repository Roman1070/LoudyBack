package app

import (
	"log/slog"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/content"
	"loudy-back/internal/services/content"
	"loudy-back/internal/storage/postgre"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	storage, err := postgre.New()
	if err != nil {
		panic(err)
	}

	contentService := content.New(log, storage)

	grpcApp := grpcApp.New(log, contentService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
