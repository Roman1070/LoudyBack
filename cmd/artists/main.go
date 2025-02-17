package main

import (
	common "loudy-back/cmd"
	appArtists "loudy-back/internal/app/artists"
	"loudy-back/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	artistsApp, err := appArtists.New(log, cfg.GRPC.Artists.Port)
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
