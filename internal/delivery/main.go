package main

import (
	"loudy-back/internal/config"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.MustLoad()
	authClient, _ := NewAuthClient(cfg.Clients.Auth.Address, cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)

	router := mux.NewRouter()

	router.HandleFunc("/api/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)
}
