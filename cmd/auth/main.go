package main

import (
	common "loudy-back/cmd"
	appAuth "loudy-back/internal/app/auth"
	"loudy-back/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	authApp := appAuth.New(log, cfg.GRPC.Auth.Port, cfg.TokenTTL)

	defer func() {
		if r := recover(); r != nil {
			log.Info("Recovered from panic:", r)
		}
	}()

	go func() {
		authApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	authApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
