package albums

import (
	"context"
	"fmt"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/albums"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Albums interface {
	Album(ctx context.Context, id primitive.ObjectID) (models.Album, error)
	CreateAlbum(ctx context.Context, name, cover string, releaseDate time.Time, artistsIds []primitive.ObjectID) (*emptypb.Empty, error)
}

type AlbumsClient struct {
	log              *slog.Logger
	ALbumsGRPCClient albumsv1.AlbumsClient
}

func NewAlbumsClient(addr string, timeout time.Duration, retriesCount int, log *slog.Logger) (*AlbumsClient, error) {
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
		ALbumsGRPCClient: albumsv1.NewAlbumsClient(cc),
	}, nil
}
