package main

import (
	common "loudy-back/cmd"
	artistsv1 "loudy-back/gen/go/artists"
	appAlbums "loudy-back/internal/app/albums"
	"loudy-back/internal/config"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)
	cc, err := grpc.NewClient(common.GrpcArtistsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor())

	if err != nil {
		panic(err)
	}

	artistsClient := artistsv1.NewArtistsClient(cc)

	albumsApp, err := appAlbums.New(log, cfg.GRPC.Albums.Port, artistsClient)
	if err != nil {
		panic(err)
	}

	go func() {
		albumsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	albumsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
