package main

import (
	"fmt"
	common "loudy-back/cmd"
	mongo_db "loudy-back/configs/mongo"
	artistsv1 "loudy-back/gen/go/artists"
	"loudy-back/internal/config"
	"loudy-back/internal/middlewares"
	repositoryAlbums "loudy-back/internal/storage/albums"
	repositoryArtists "loudy-back/internal/storage/artists"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	cc, err := grpc.NewClient(common.GrpcArtistsAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor())
	if err != nil {
		panic(err)
	}

	artistsGRPCClient := artistsv1.NewArtistsClient(cc)

	mongoDb, err := mongo_db.Connect()
	if err != nil {
		return
	}

	albumsRepo := repositoryAlbums.NewStorage(mongoDb, "albums", artistsGRPCClient, log)
	artistsRepo := repositoryArtists.NewStorage(mongoDb, "artists", log, albumsRepo)

	// postgreStorage, err := postgre.New()
	// if err != nil {
	// 	panic("can't create db connection: " + err.Error())
	// }

	authClient, _ := NewAuthClient(common.GrpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	artistsClient, _ := NewArtistsClient(common.GrpcArtistsAddress(cfg), cfg.Clients.Artists.Timeout, cfg.Clients.Artists.RetriesCount, artistsRepo)
	albumsClient, _ := NewAlbumsClient(common.GrpcAlbumsAddress(cfg), cfg.Clients.Albums.Timeout, cfg.Clients.Albums.RetriesCount, log, albumsRepo)

	router := mux.NewRouter()

	router.HandleFunc("/api/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/artist", artistsClient.Artist).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/artist", artistsClient.CreateArtist).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/album", albumsClient.Album).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/album", albumsClient.CreateAlbum).Methods(http.MethodPost, http.MethodOptions)

	handler := middlewares.CorsMiddleware(router)
	fmt.Println("Server is listening...")

	http.ListenAndServe(os.Getenv(AppPortEnv), handler)
}
