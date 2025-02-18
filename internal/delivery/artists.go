package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/storage"
	"loudy-back/utils"
	"net/http"
	"strings"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistsProvider interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
	CreateArtist(ctx context.Context, name, bio, cover string) (*emptypb.Empty, error)
}

type ArtistsClient struct {
	ArtistsGRPCClient artistsv1.ArtistsClient
}

func NewArtistsClient(addr string, timeout time.Duration, retriesCount int) (*ArtistsClient, error) {
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))
	if err != nil {
		slog.Error("[NewArtistsClient] client error: " + err.Error())
		return nil, fmt.Errorf("%s", "[NewArtistsClient] client  error: "+err.Error())
	}

	return &ArtistsClient{
		ArtistsGRPCClient: artistsv1.NewArtistsClient(cc),
	}, nil
}

func (c *ArtistsClient) CreateArtist(w http.ResponseWriter, r *http.Request) {
	slog.Info("[CreateArtist] client started ")

	var request models.CreateArtistRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	_, err = c.ArtistsGRPCClient.CreateArtist(r.Context(), request.ToGRPC())
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		if strings.Contains(err.Error(), storage.ErrArtistAlreadyExists.Error()) {
			utils.WriteError(w, fmt.Sprintf("Artist %v already exists.", request.Name))
		} else {
			utils.WriteError(w, "Internal error")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *ArtistsClient) Artist(w http.ResponseWriter, r *http.Request) {
	slog.Info("[Artist] client started")

	id := r.URL.Query().Get("id")

	artist, err := c.ArtistsGRPCClient.Artist(r.Context(), &artistsv1.ArtistRequest{
		Id: id,
	})
	if err != nil {
		slog.Error("[Artist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(artist)
	if err != nil {
		slog.Error("[Artist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
