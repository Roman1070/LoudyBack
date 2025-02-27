package main

import (
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	artistsv1 "loudy-back/gen/go/artists"
	appAlbums "loudy-back/internal/app/albums"
	"loudy-back/internal/config"
	repositoryAlbums "loudy-back/internal/storage/albums"
	repositoryArtists "loudy-back/internal/storage/artists"
	repositoryTracks "loudy-back/internal/storage/tracks"
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

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return
	}

	albumsRepo := repositoryAlbums.NewStorage(mongoDb, "albums", artistsClient, log)
	artistsRepo := repositoryArtists.NewStorage(mongoDb, "artists", log, albumsRepo)
	tracksRepo := repositoryTracks.NewStorage(mongoDb, "tracks", log)

	albumsApp, err := appAlbums.New(log, cfg.GRPC.Albums.Port, artistsClient, artistsRepo, tracksRepo)
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
