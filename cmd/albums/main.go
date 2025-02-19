package main

import (
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	albumsv1 "loudy-back/gen/go/albums"
	artistsv1 "loudy-back/gen/go/artists"
	appAlbums "loudy-back/internal/app/albums"
	"loudy-back/internal/config"
	repositoryArtists "loudy-back/internal/storage/artists"
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

	cc, err = grpc.NewClient(common.GrpcAlbumsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor())

	if err != nil {
		panic(err)
	}

	albumsClient := albumsv1.NewAlbumsClient(cc)

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		panic(err)
	}

	repo := repositoryArtists.NewStorage(mongoDb, "artists", log, albumsClient)

	albumsApp, err := appAlbums.New(log, cfg.GRPC.Albums.Port, artistsClient, repo)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Info("Recovered from panic:", r)
		}
	}()

	go func() {
		albumsApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	albumsApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
