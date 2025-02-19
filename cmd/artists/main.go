package main

import (
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	albumsv1 "loudy-back/gen/go/albums"
	artistsv1 "loudy-back/gen/go/artists"
	appArtists "loudy-back/internal/app/artists"
	"loudy-back/internal/config"
	repositoryAlbums "loudy-back/internal/storage/albums"
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

	cc, err = grpc.NewClient(common.GrpcAlbumsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor())

	if err != nil {
		panic(err)
	}

	artistsClient := artistsv1.NewArtistsClient(cc)

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		panic(err)
	}

	albumsRepo := repositoryAlbums.NewStorage(mongoDb, "albums", artistsClient, log)

	artistsApp, err := appArtists.New(log, cfg.GRPC.Artists.Port, albumsClient, albumsRepo)
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
