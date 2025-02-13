package main

import (
	common "loudy-back/cmd"
	appContent "loudy-back/internal/app/content"
	"loudy-back/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := common.SetupLogger(cfg.Env)

	contentApp := appContent.New(log, cfg.GRPC.Content.Port)

	go func() {
		contentApp.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	contentApp.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
