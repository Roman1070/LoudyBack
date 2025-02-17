package content

import (
	models "loudy-back/internal/domain/models/content"
	"time"
)

type dtoArtist struct {
	Name       string     `bson:"name"`
	Cover      string     `bson:"cover"`
	Bio        string     `bson:"bio"`
	LikesCount uint32     `bson:"likes_count"`
	Albums     []dtoAlbum `bson:"albums"`
}

type dtoAlbum struct {
	Name        string           `bson:"name"`
	Artists     []dtoArtistLight `bson:"artists"`
	Cover       string           `bson:"cover"`
	ReleaseDate time.Time        `bson:"release_date"`
	Tracks      []dtoTrack       `bson:"tracks"`
}
type dtoArtistLight struct {
	Name  string `bson:"name"`
	Cover string `bson:"cover"`
}
type dtoTrack struct {
	Name    string `bson:"name"`
	AlbumId uint32 `bson:"album_id"`
}

func (artist *dtoArtist) toCommonModel() models.Artist {
	albums := make([]models.AlbumLight, len(artist.Albums))

	for i, album := range artist.Albums {
		tracks := make([]models.TrackLight, len(artist.Albums[i].Tracks))
		for j, track := range artist.Albums[i].Tracks {
			tracks[j] = models.TrackLight{
				Name:    track.Name,
				AlbumId: track.AlbumId,
			}
		}
		albums[i] = models.AlbumLight{
			Name:  album.Name,
			Cover: album.Cover,
			Year:  uint32(album.ReleaseDate.Year()),
		}

	}

	return models.Artist{
		Name:       artist.Name,
		Cover:      artist.Cover,
		Bio:        artist.Bio,
		LikesCount: uint32(artist.LikesCount),
		Albums:     albums,
	}
}
