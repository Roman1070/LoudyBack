package grpcApp

import (
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "loudy-back/internal/app"
	artistsgrpc "loudy-back/internal/grpc/artists"
)

// New creates new gRPC server app.
func New(
	log *slog.Logger,
	artistsService artistsgrpc.Artists,
	port int,
) *common.App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			//logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
		// Add any other option (check functions starting with logging.With).
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(common.InterceptorLogger(log), loggingOpts...),
	))

	artistsgrpc.Register(gRPCServer, artistsService, log)

	return &common.App{
		Log:        log,
		GRPCServer: gRPCServer,
		Port:       port,
	}
}
