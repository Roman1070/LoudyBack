package main

import (
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	appAlbums "loudy-back/internal/app/albums"
	"loudy-back/internal/config"
	"loudy-back/internal/delivery/artists"
	repositoryArtists "loudy-back/internal/storage/artists"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)
	mongoDb, err := mongo_db.Connect()
	if err != nil {
		panic(err)
	}

	artistsStorage := repositoryArtists.NewStorage(mongoDb, "artists", log)

	artistsClient, _ := artists.NewArtistsClient(common.GrpcArtistsAddress(cfg),
		cfg.Clients.Artists.Timeout, cfg.Clients.Artists.RetriesCount, artistsStorage)

	albumsApp, err := appAlbums.New(log, cfg.GRPC.Albums.Port, artistsClient.ArtistsGRPCClient)
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
