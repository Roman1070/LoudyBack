package content

import (
	"log/slog"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentService struct {
	log            *slog.Logger
	contentCreator ContentCreator
	tokenTTL       time.Duration
}

type ContentCreator interface {
	CreateArtist(name, cover, bio string) (uint32, error)
	CreateTrack(name, file string, albumId uint32) (*emptypb.Empty, error)
	CreateAlbum(artistsIds, tracksIds []uint32, name, cover string, releaseDate time.Time) (uint32, error)
}

func New(log *slog.Logger, contentCreator ContentCreator, tokenTTL time.Duration) *ContentService {
	return &ContentService{
		contentCreator: contentCreator,
		log:            log,
		tokenTTL:       tokenTTL,
	}
}
