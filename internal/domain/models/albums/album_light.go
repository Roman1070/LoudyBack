package models

import (
	"errors"
	"log/slog"
	albumsv1 "loudy-back/gen/go/albums"
	models "loudy-back/internal/domain/models/artists"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlbumLight struct {
	ID          primitive.ObjectID
	Name        string
	Cover       string
	ReleaseDate string
}

func AlbumsLightToGRPC(albums []AlbumLight) *albumsv1.AlbumsLightResponse {
	result := make([]*albumsv1.AlbumLight, len(albums))

	for i, album := range albums {
		result[i] = &albumsv1.AlbumLight{
			Id:          album.ID.Hex(),
			Name:        album.Name,
			Cover:       album.Cover,
			ReleaseDate: album.ReleaseDate,
		}
	}

	return &albumsv1.AlbumsLightResponse{
		Albums: result,
	}
}

func GRPCResponseToAlbumsLight(albums []*albumsv1.AlbumLight) ([]models.AlbumLight, error) {
	result := make([]models.AlbumLight, len(albums))

	for i, album := range albums {
		id, err := primitive.ObjectIDFromHex(album.Id)
		if err != nil {
			slog.Error("[GRPCResponseToAlbumsLight] error: " + err.Error())
			return nil, errors.New("[GRPCResponseToAlbumsLight] error: " + err.Error())
		}

		result[i] = models.AlbumLight{
			ID:          id,
			Name:        album.Name,
			Cover:       album.Cover,
			ReleaseDate: album.ReleaseDate,
		}
	}

	return result, nil
}
