package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/albums"
	"loudy-back/utils"
	"net/http"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type AlbumsProvider interface {
	AlbumsLight(ctx context.Context, ids []primitive.ObjectID) ([]models.AlbumLight, error)
}

type AlbumsClient struct {
	log              *slog.Logger
	albumsProvider   AlbumsProvider
	AlbumsGRPCClient albumsv1.AlbumsClient
}

func NewAlbumsClient(addr string, timeout time.Duration, retriesCount int, log *slog.Logger, albumsProvider AlbumsProvider) (*AlbumsClient, error) {
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))

	if err != nil {
		slog.Error("[NewartistsClient] client error: " + err.Error())
		return nil, fmt.Errorf("%s", "[NewartistsClient] client  error: "+err.Error())
	}

	return &AlbumsClient{
		log:              log,
		albumsProvider:   albumsProvider,
		AlbumsGRPCClient: albumsv1.NewAlbumsClient(cc),
	}, nil
}

func (c *AlbumsClient) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	c.log.Info("[CreateAlbum] client started")

	var request models.CreateAlbumRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.log.Error("[CreateAlbum] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	_, err = c.AlbumsGRPCClient.CreateAlbum(r.Context(), &albumsv1.CreateAlbumRequest{
		Name:        request.Name,
		ArtistsIds:  request.ArtistsIds,
		Cover:       request.Cover,
		ReleaseDate: request.ReleaseDate,
	})

	if err != nil {
		c.log.Error("[CreateAlbum] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
	}

	w.WriteHeader(http.StatusOK)
}

func (c *AlbumsClient) Album(w http.ResponseWriter, r *http.Request) {
	c.log.Info("[Album] client started")

	resp, err := c.AlbumsGRPCClient.Album(r.Context(), &albumsv1.AlbumRequest{
		Id: r.URL.Query().Get("id"),
	})
	if err != nil {
		c.log.Error("[Album] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		c.log.Error("[Album] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
