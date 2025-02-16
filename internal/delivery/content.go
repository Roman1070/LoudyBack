package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	contentv1 "loudy-back/gen/go/content"
	models "loudy-back/internal/domain/models/content"
	"loudy-back/utils"
	"net/http"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ContentClient struct {
	contentProvider ContentProvider
	contentCreator  contentv1.ContentClient
}

type ContentProvider interface {
	Artist(ctx context.Context, name string) (models.Artist, error)
	Album(ctx context.Context, id uint32) (models.Album, error)
	SearchContent(ctx context.Context, input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error)
}

func NewContentClient(addr string, timeout time.Duration, retriesCount int, contentProvider ContentProvider) (*ContentClient, error) {
	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))
	if err != nil {
		slog.Error("[NewContentClient] client error: " + err.Error())
		return nil, fmt.Errorf("%s", "[NewContentClient] client  error: "+err.Error())
	}

	return &ContentClient{
		contentProvider: contentProvider,
		contentCreator:  contentv1.NewContentClient(cc),
	}, nil
}

func (c *ContentClient) Artist(w http.ResponseWriter, r *http.Request) {
	slog.Info("client start [Artist]")
	name := r.URL.Query().Get("name")

	artist, err := c.contentProvider.Artist(r.Context(), name)
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

func (c *ContentClient) CreateArtist(w http.ResponseWriter, r *http.Request) {
	slog.Info("client start [CreateArtist]")

	var request models.CreateArtistRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	resp, err := c.contentCreator.CreateArtist(r.Context(), request.ToGRPC())
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
