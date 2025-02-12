package app

import (
	"log/slog"
	common "loudy-back/internal/app"
	grpcApp "loudy-back/internal/app/grpc/auth"
	"loudy-back/internal/services/auth"
	"loudy-back/internal/storage/postgre"
	"time"
)

type App struct {
	GRPCServer *common.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
) *App {
	storage, err := postgre.New()
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, tokenTTL)

	grpcApp := grpcApp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
