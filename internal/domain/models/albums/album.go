package models

import (
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Album struct {
	ID          primitive.ObjectID `bson:"omitempty,_id"`
	Name        string
	Cover       string
	ReleaseDate time.Time                `bson:"release_date"`
	Artists     []models.ArtistLight     `bson:"artists"`
	Tracks      []trackModels.TrackLight `bson:"tracks"`
}

func (a *Album) ToGRPC() *albumsv1.AlbumResponse {
	artists := make([]*albumsv1.ArtistLight, len(a.Artists))

	for i, artist := range a.Artists {
		artists[i] = &albumsv1.ArtistLight{
			Id:   artist.ID.Hex(),
			Name: artist.Name,
		}
	}
	tracks := make([]*albumsv1.TrackLight, len(a.Tracks))

	for i, track := range a.Tracks {
		tracks[i] = &albumsv1.TrackLight{
			Id:   track.ID.Hex(),
			Name: track.Name,
		}
	}
	return &albumsv1.AlbumResponse{
		Id:          a.ID.Hex(),
		Name:        a.Name,
		Cover:       a.Cover,
		ReleaseDate: timestamppb.New(a.ReleaseDate),
		Artists:     artists,
		Tracks:      tracks,
	}
}
