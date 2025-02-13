package content

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentService struct {
	log            *slog.Logger
	contentCreator ContentCreator
}

type ContentCreator interface {
	CreateArtist(ctx context.Context, name, cover, bio string) (uint32, error)
	CreateTrack(ctx context.Context, name, file string, albumId uint32) (*emptypb.Empty, error)
	CreateAlbum(ctx context.Context, name, cover string, tracksIds []uint32, releaseDate time.Time) (uint32, error)
}

func New(log *slog.Logger, contentCreator ContentCreator) *ContentService {
	return &ContentService{
		contentCreator: contentCreator,
		log:            log,
	}
}
