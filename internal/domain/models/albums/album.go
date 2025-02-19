package models

import (
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	ReleaseDate string
	Artists     []models.ArtistLight
	Tracks      []trackModels.TrackLight
}

type AlbumPreliminary struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	ReleaseDate string
	ArtistsIds  []primitive.ObjectID
	TracksIds   []primitive.ObjectID
}

func (a *Album) ToGRPC() *albumsv1.AlbumData {
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

	return &albumsv1.AlbumData{
		Id:          a.ID.Hex(),
		Name:        a.Name,
		Cover:       a.Cover,
		ReleaseDate: a.ReleaseDate,
		Artsits:     artists,
		Tracks:      tracks,
	}
}
