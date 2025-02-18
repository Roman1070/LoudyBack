package models

import (
	"errors"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artist struct {
	ID         primitive.ObjectID
	Name       string
	Albums     []AlbumLight
	Cover      string
	Bio        string
	LikesCount uint32 `bson:"likes_count"`
}

func (artist *Artist) ToGRPC() *artistsv1.ArtistResponse {
	albums := make([]*artistsv1.AlbumLight, len(artist.Albums))

	for i, album := range artist.Albums {
		albums[i] = &artistsv1.AlbumLight{
			Id:          album.ID.Hex(),
			Cover:       album.Cover,
			Name:        album.Name,
			ReleaseDate: album.ReleaseDate,
		}
	}

	return &artistsv1.ArtistResponse{
		Artist: &artistsv1.ArtistData{
			Id:         artist.ID.Hex(),
			Name:       artist.Name,
			Bio:        artist.Bio,
			Cover:      artist.Cover,
			LikesCount: artist.LikesCount,
			Albums:     albums,
		},
	}
}

func ModelsFromArtistDataLight(artists []*artistsv1.ArtistLight) ([]ArtistLight, error) {
	result := make([]ArtistLight, len(artists))

	for i, artist := range artists {
		id, err := primitive.ObjectIDFromHex(artist.Id)
		if err != nil {
			slog.Error("[ModelsFromArtistData] error: " + err.Error())
			return nil, errors.New("[ModelsFromArtistData] error: " + err.Error())
		}

		result[i] = ArtistLight{
			ID:   id,
			Name: artist.Name,
		}
	}

	return result, nil
}

func ArtistsLightToGRPC(artists []ArtistLight) *artistsv1.ArtistsLightResponse {
	result := make([]*artistsv1.ArtistLight, len(artists))

	for i, artist := range artists {
		result[i] = artist.ToGRPC()
	}

	return &artistsv1.ArtistsLightResponse{
		Artists: result,
	}
}
