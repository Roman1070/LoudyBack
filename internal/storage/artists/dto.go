package artists

import (
	modelsAlbums "loudy-back/internal/domain/models/albums"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dtoArtist struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Name       string               `bson:"name"`
	Cover      string               `bson:"cover"`
	Bio        string               `bson:"bio"`
	LikesCount uint32               `bson:"likes_count"`
	AlbumsIds  []primitive.ObjectID `bson:"albums_ids"`
}
type dtoArtistLight struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func (artist *dtoArtist) toCommonModel(albums []modelsAlbums.AlbumLight) models.Artist {
	albumsInResult := make([]models.AlbumLight, len(albums))

	for i, album := range albums {
		albumsInResult[i] = models.AlbumLight{
			ID:          album.ID,
			Name:        album.Name,
			Cover:       album.Cover,
			ReleaseDate: album.ReleaseDate,
		}
	}
	return models.Artist{
		ID:         artist.ID,
		Name:       artist.Name,
		Cover:      artist.Cover,
		Bio:        artist.Bio,
		LikesCount: uint32(artist.LikesCount),
		Albums:     albumsInResult,
	}
}

func toCommonModels(artists []dtoArtistLight) []models.ArtistLight {
	result := make([]models.ArtistLight, len(artists))

	for i, artist := range artists {
		result[i] = models.ArtistLight{
			ID:   artist.ID,
			Name: artist.Name,
		}
	}
	return result
}
