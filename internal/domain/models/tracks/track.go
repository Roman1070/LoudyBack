package models

import (
	tracksv1 "loudy-back/gen/go/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Track struct {
	ID       primitive.ObjectID
	Name     string
	AlbumID  primitive.ObjectID
	Cover    string
	Artists  []ArtistLight
	Duration uint16
}

type TrackPreliminary struct {
	ID         primitive.ObjectID
	Name       string
	AlbumID    primitive.ObjectID
	ArtistsIds []primitive.ObjectID
	Duration   uint16
}

type ArtistLight struct {
	ID   primitive.ObjectID
	Name string
}

func (t *Track) ToGRPC() *tracksv1.TrackData {
	artists := make([]*tracksv1.ArtistLight, len(t.Artists))

	for i, artist := range t.Artists {
		artists[i] = &tracksv1.ArtistLight{
			Id:   artist.ID.Hex(),
			Name: artist.Name,
		}
	}

	return &tracksv1.TrackData{
		Id:       t.ID.Hex(),
		Name:     t.Name,
		AlbumId:  t.AlbumID.Hex(),
		Cover:    t.Cover,
		Duration: uint32(t.Duration),
		Artists:  artists,
	}
}
