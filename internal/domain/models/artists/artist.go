package models

import (
	"errors"
	"log/slog"
	artistsv1 "loudy-back/gen/go/artists"
	"loudy-back/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artist struct {
	ID         primitive.ObjectID
	Name       string
	AlbumsIds  []primitive.ObjectID `bson:"albums_ids"`
	Cover      string
	Bio        string
	LikesCount uint32 `bson:"likes_count"`
}

func (artist *Artist) ToGRPC() *artistsv1.ArtistResponse {
	ids := utils.IdsToStringArray(artist.AlbumsIds)

	return &artistsv1.ArtistResponse{
		Artist: &artistsv1.ArtistData{
			Id:         artist.ID.Hex(),
			Name:       artist.Name,
			Bio:        artist.Bio,
			Cover:      artist.Cover,
			LikesCount: artist.LikesCount,
			AlbumsIds:  ids,
		},
	}
}

func ModelsFromArtistData(artists []*artistsv1.ArtistData) ([]Artist, error) {
	result := make([]Artist, len(artists))

	for i, artist := range artists {
		albumsIds, err := utils.StringsToIdsArray(artist.AlbumsIds)

		if err != nil {
			slog.Error("[ModelsFromArtistData] error: " + err.Error())
			return nil, errors.New("[ModelsFromArtistData] error: " + err.Error())
		}

		id, err := primitive.ObjectIDFromHex(artist.Id)
		if err != nil {
			slog.Error("[ModelsFromArtistData] error: " + err.Error())
			return nil, errors.New("[ModelsFromArtistData] error: " + err.Error())
		}

		result[i] = Artist{
			ID:         id,
			Name:       artist.Name,
			Bio:        artist.Bio,
			Cover:      artist.Cover,
			LikesCount: artist.LikesCount,
			AlbumsIds:  albumsIds,
		}
	}
	return result, nil
}

func ArtistsToGRPC(artists []Artist) *artistsv1.ArtistsResponse {
	result := make([]*artistsv1.ArtistData, len(artists))

	for i, artist := range artists {
		ids := make([]string, len(artist.AlbumsIds))
		for j, id := range artist.AlbumsIds {
			ids[j] = id.Hex()
		}

		result[i] = &artistsv1.ArtistData{
			Id:         artist.ID.Hex(),
			Name:       artist.Name,
			Bio:        artist.Bio,
			Cover:      artist.Cover,
			LikesCount: artist.LikesCount,
			AlbumsIds:  ids,
		}
	}

	return &artistsv1.ArtistsResponse{
		Artists: result,
	}
}
