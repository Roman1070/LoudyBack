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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type ArtistsClient struct {
	artistsProvider   ArtistProvider
	ArtistsGRPCClient artistsv1.ArtistsClient
}

type ArtistProvider interface {
	Artist(ctx context.Context, id primitive.ObjectID) (models.Artist, error)
	ArtistsLight(ctx context.Context, id []primitive.ObjectID) ([]models.ArtistLight, error)
	ArtistByName(ctx context.Context, name string) (models.Artist, error)
	ArtistLightByName(ctx context.Context, name string) (models.ArtistLight, error)
}

func NewArtistsClient(addr string, timeout time.Duration, retriesCount int, artistsProvider ArtistProvider) (*ArtistsClient, error) {
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
		artistsProvider:   artistsProvider,
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

	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")

	var artist interface{}
	var err error
	if len(name) == 0 {
		var id primitive.ObjectID
		id, err = primitive.ObjectIDFromHex(idStr)
		if err != nil {
			slog.Error("[Artist] client error: " + err.Error())
			utils.WriteError(w, "Internal error")
			return
		}
		artist, err = c.artistsProvider.Artist(r.Context(), id)

	} else {
		artist, err = c.artistsProvider.ArtistByName(r.Context(), name)
	}

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
