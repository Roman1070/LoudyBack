package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	contentv1 "loudy-back/gen/go/content"
	models "loudy-back/internal/domain/models/content"
	"loudy-back/internal/services/content"
	"loudy-back/internal/storage"
	"loudy-back/utils"
	"net/http"
	"strings"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ContentClient struct {
	contentProvider content.ContentProvider
	contentCreator  contentv1.ContentClient
}

func NewContentClient(addr string, timeout time.Duration, retriesCount int, contentProvider content.ContentProvider) (*ContentClient, error) {
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
	slog.Info("[CreateArtist] client started ")

	var request models.CreateArtistRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("[CreateArtist] client error: " + err.Error())
		utils.WriteError(w, "Internal error")
		return
	}

	_, err = c.contentCreator.CreateArtist(r.Context(), request.ToGRPC())
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
