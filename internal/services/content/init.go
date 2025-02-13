package content

import (
	"log/slog"
	"loudy-back/internal/domain/models"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentService struct {
	log             *slog.Logger
	contentProvider ContentProvider
	tokenTTL        time.Duration
}

type ContentProvider interface {
	CreateArtist(name, cover, bio string) (uint32, error)
	Artist(name string) (models.Artist, error)
	CreateTrack(name, file string, albumId uint32) (*emptypb.Empty, error)
	CreateAlbum(artistsIds, tracksIds []uint32, name, cover string, releaseDate time.Time) (uint32, error)
	Album(id uint32) (models.Album, error)
	SearchContent(input string) ([]models.ArtistLight, []models.AlbumLight, []models.TrackLight, error)
}

/*

message CreateAlbumRequest{
    repeated uint32 artists_ids=1;
    string name=2;
    string cover = 3;
    repeated uint32 tracks_ids=4;
    google.protobuf.Timestamp release_date = 5;
}
*/
// New returns a new instance of the Auth service.
func New(log *slog.Logger, contentProvider ContentProvider, tokenTTL time.Duration) *ContentService {
	return &ContentService{
		contentProvider: contentProvider,
		log:             log,
		tokenTTL:        tokenTTL,
	}
}
