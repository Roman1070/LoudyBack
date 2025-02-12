package postgre

import (
	"fmt"
	"loudy-back/configs/postgres"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

const emptyValue = -1

func New() (*Storage, error) {
	// cfg := config.MustLoad()

	// retryOptions := []grpcretry.CallOption{
	// 	grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
	// 	grpcretry.WithMax(uint(5)),
	// 	grpcretry.WithPerRetryTimeout(5 * time.Second),
	// }

	pool, err := postgres.LoadPgxPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// profilesClient, err := grpc.NewClient(common.GrpcProfilesAddress(cfg),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
	// 		grpcretry.UnaryClientInterceptor(retryOptions...),
	// 	))

	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }

	return &Storage{db: pool}, nil
}
