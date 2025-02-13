package main

import (
	"loudy-back/internal/config"
	"loudy-back/internal/storage/postgre"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.MustLoad()

	storage, err := postgre.New()
	if err != nil {
		panic("can't create db connection: " + err.Error())
	}

	authClient, _ := NewAuthClient(cfg.Clients.Auth.Address, cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	contentClient, _ := NewContentClient(cfg.Clients.Content.Address, cfg.Clients.Content.Timeout, cfg.Clients.Content.RetriesCount, storage)

	router := mux.NewRouter()

	router.HandleFunc("/api/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/artist", contentClient.Artist).Methods(http.MethodPost, http.MethodOptions)
}
