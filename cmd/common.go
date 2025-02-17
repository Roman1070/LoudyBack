package common

import (
	"fmt"
	"log/slog"
	"loudy-back/internal/config"
	"loudy-back/internal/lib/logger/handlers/slogpretty"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = SetupPrettySlog()
	case EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func SetupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func GrpcAuthAddress(cfg *config.Config) string {
	return fmt.Sprintf("auth-go:%v", cfg.GRPC.Auth.Port)
}

func GrpcArtistsAddress(cfg *config.Config) string {
	return fmt.Sprintf("artists-go:%v", cfg.GRPC.Artists.Port)
}

func GrpcAlbumsAddress(cfg *config.Config) string {
	return fmt.Sprintf("albums-go:%v", cfg.GRPC.Albums.Port)
}
