package models

import (
	playlistsv1 "loudy-back/gen/go/playlists"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	CreatorID   primitive.ObjectID
	CreatorName string
	Tracks      []models.Track
}

type PlaylistPreliminary struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	CreatorID   primitive.ObjectID
	CreatorName string
	TracksIds   []primitive.ObjectID
}

func (p *Playlist) ToGRPC() *playlistsv1.PlaylistData {
	tracks := make([]*playlistsv1.TrackData, len(p.Tracks))

	for i, track := range p.Tracks {
		artists := make([]*playlistsv1.ArtistLight, len(track.Artists))

		for j, artist := range track.Artists {
			artists[j] = &playlistsv1.ArtistLight{
				Id:   artist.ID.Hex(),
				Name: artist.Name,
			}
		}

		tracks[i] = &playlistsv1.TrackData{
			Id:       track.ID.Hex(),
			Name:     track.Name,
			Filename: track.Filename,
			Cover:    track.Cover,
			AlbumId:  track.AlbumID.Hex(),
			Artists:  artists,
			Duration: uint32(track.Duration),
		}
	}

	return &playlistsv1.PlaylistData{
		Id:          p.ID.Hex(),
		Name:        p.Name,
		Cover:       p.Cover,
		CreatorId:   p.CreatorID.Hex(),
		CreatorName: p.CreatorName,
		Tracks:      tracks,
	}
}
