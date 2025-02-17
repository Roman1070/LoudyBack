package models

import (
	artistsv1 "loudy-back/gen/go/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artist struct {
	Name       string
	Albums     []primitive.ObjectID
	Cover      string
	Bio        string
	LikesCount uint32 `bson:"likes_count"`
}

func (artist *Artist) ToGRPC() *artistsv1.ArtistResponse {

	ids := make([]string, len(artist.Albums))
	for i, id := range artist.Albums {
		ids[i] = id.Hex()
	}

	return &artistsv1.ArtistResponse{
		Artist: &artistsv1.ArtistData{
			Name:       artist.Name,
			Bio:        artist.Bio,
			Cover:      artist.Cover,
			LikesCount: artist.LikesCount,
			AlbumsIds:  ids,
		},
	}
}
