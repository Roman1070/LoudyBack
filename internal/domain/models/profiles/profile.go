package models

import (
	profilesv1 "loudy-back/gen/go/profiles"
	albumsModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	playlistModels "loudy-back/internal/domain/models/playlists"
	trackModels "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID             primitive.ObjectID
	UserId         uint32
	Name           string
	Avatar         string
	Bio            string
	LikesCount     uint32
	SavedTracks    []trackModels.Track
	SavedAlbums    []albumsModels.AlbumLight
	SavedArtists   []artistsModels.ArtistLight
	SavedPlaylists []playlistModels.PlaylistLight
}
type ProfilePreliminary struct {
	ID                primitive.ObjectID
	UserId            uint32
	Name              string
	Avatar            string
	Bio               string
	LikesCount        uint32
	SavedTracksIds    []primitive.ObjectID
	SavedAlbumsIds    []primitive.ObjectID
	SavedArtistsIds   []primitive.ObjectID
	SavedPlaylistsIds []primitive.ObjectID
}

func (p *Profile) ToGRPC() *profilesv1.ProfileData {
	tracks := make([]profilesv1.TrackLight, len(p.SavedTracks))

	for i, track := range p.SavedTracks {
		artists := make([]*profilesv1.ArtistLight, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j] = &profilesv1.ArtistLight{
				Id:   artist.ID.Hex(),
				Name: artist.Name,
			}
		}
		tracks[i] = profilesv1.TrackLight{
			Id:       track.ID.Hex(),
			Name:     track.Name,
			Artists:  artists,
			Cover:    track.Cover,
			AlbumId:  track.AlbumID.Hex(),
			Duration: uint32(track.Duration),
		}
	}

	artists := make([]profilesv1.ArtistLight, len(p.SavedArtists))

	for i, artist := range p.SavedArtists {
		artists[i] = profilesv1.ArtistLight{
			Id:    artist.ID.Hex(),
			Name:  artist.Name,
			Cover: artist.Cover,
		}
	}

	albums := make([]profilesv1.AlbumLight, len(p.SavedAlbums))

	for i, album := range p.SavedAlbums {
		albums[i] = profilesv1.AlbumLight{
			Id:   album.ID.Hex(),
			Name: album.Name,
		}
	}

	return &profilesv1.ProfileData{}
}
