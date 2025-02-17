package main

import (
	"fmt"
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	"loudy-back/internal/config"
	"loudy-back/internal/delivery/artists"
	"loudy-back/internal/middlewares"
	repositoryArtists "loudy-back/internal/storage/artists"
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
	log := common.SetupLogger(cfg.Env)

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return
	}

	artistsStorage := repositoryArtists.NewStorage(mongoDb, "artists", log)

	// postgreStorage, err := postgre.New()
	// if err != nil {
	// 	panic("can't create db connection: " + err.Error())
	// }

	authClient, _ := NewAuthClient(common.GrpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	artistsClient, _ := artists.NewArtistsClient(common.GrpcArtistsddress(cfg), cfg.Clients.Artists.Timeout, cfg.Clients.Artists.RetriesCount, artistsStorage)

	router := mux.NewRouter()

	router.HandleFunc("/api/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/artist", artistsClient.Artist).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/artist", artistsClient.CreateArtist).Methods(http.MethodPost, http.MethodOptions)

	handler := middlewares.CorsMiddleware(router)
	fmt.Println("Server is listening...")

	http.ListenAndServe(os.Getenv(AppPortEnv), handler)
}
