package main

import (
	common "loudy-back/cmd"
	albumsv1 "loudy-back/gen/go/albums"
	appArtists "loudy-back/internal/app/artists"
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

	cc, err := grpc.NewClient(common.GrpcAlbumsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor())

	if err != nil {
		panic(err)
	}

	albumsClient := albumsv1.NewAlbumsClient(cc)

	artistsApp, err := appArtists.New(log, cfg.GRPC.Artists.Port, albumsClient)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Info("Recovered from panic:", r)
		}
	}()
	go func() {
		artistsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	artistsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
