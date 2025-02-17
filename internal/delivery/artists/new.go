package artists

import (
	"context"
	"fmt"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	models "loudy-back/internal/domain/models/artists"
	"loudy-back/internal/services/artists"
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
	artistsProvider   artists.Artists
	ArtistsGRPCClient artistsv1.ArtistsClient
}

func NewArtistsClient(addr string, timeout time.Duration, retriesCount int, artists artists.Artists) (*ArtistsClient, error) {
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

	return &ArtistsClient{
		artistsProvider:   artists,
		ArtistsGRPCClient: artistsv1.NewArtistsClient(cc),
	}, nil
}
