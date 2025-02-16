package main

import (
	"fmt"
	"log"
	"loudy-back/internal/config"
	"loudy-back/internal/middlewares"
	"loudy-back/internal/storage/postgre"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	AppPortEnv      = "APP_PORT"
	AppSecretEnv    = "APP_SECRET"
	TokenCookieName = "token"
	EmptyValue      = int64(-1)
)

func main() {
	cfg := config.MustLoad()

	postgreStorage, err := postgre.New()
	if err != nil {
		panic("can't create db connection: " + err.Error())
	}

	authClient, _ := NewAuthClient(cfg.Clients.Auth.Address, cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	contentClient, _ := NewContentClient(cfg.Clients.Content.Address, cfg.Clients.Content.Timeout, cfg.Clients.Content.RetriesCount, postgreStorage)

	router := mux.NewRouter()

	router.HandleFunc("/api/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/artist", contentClient.Artist).Methods(http.MethodGet, http.MethodOptions)

	handler := middlewares.CorsMiddleware(router)
	fmt.Println("Server is listening...")

	log.Fatal(http.ListenAndServe(os.Getenv(AppPortEnv), handler))
}
