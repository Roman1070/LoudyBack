package albums

import (
	"context"
	"errors"
	"log/slog"
	models "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	trackModels "loudy-back/internal/domain/models/tracks"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *AlbumsService) Album(ctx context.Context, id primitive.ObjectID) (models.Album, error) {
	s.log.Info("[Album] service started")

	preliminaryAlbum, resultErr := s.albums.Album(ctx, id)
	if resultErr != nil {
		s.log.Error("[Album] service error: " + resultErr.Error())
		return models.Album{}, errors.New("[Album] service error: " + resultErr.Error())
	}

	wg := sync.WaitGroup{}
	artistsChan := make(chan []artistsModels.ArtistLight, 1)
	tracksChan := make(chan []trackModels.TrackLight, 1)

	errorChan := make(chan error, 2)

	wg.Add(1)
	go func(ctx context.Context, artistsProvider ArtistsProvider, ids []primitive.ObjectID, log *slog.Logger) {
		artists, err := artistsProvider.ArtistsLight(ctx, ids)
		if err != nil {
			errorChan <- errors.New("[Album] service error: " + err.Error())
			return
		}

		artistsChan <- artists
		wg.Done()

	}(ctx, s.artistsProvider, preliminaryAlbum.ArtistsIds, s.log)

	wg.Add(1)
	go func(ctx context.Context, tracksProvider TracksProvider, ids []primitive.ObjectID, log *slog.Logger) {
		tracks, err := tracksProvider.TracksLight(ctx, ids)
		if err != nil {
			errorChan <- errors.New("[Album] service error: " + err.Error())
			return
		}

		tracksChan <- tracks
		wg.Done()

	}(ctx, s.tracksProvider, preliminaryAlbum.TracksIds, s.log)

	wg.Wait()

	errorsCount := len(errorChan)
	for i := 0; i < errorsCount; i++ {
		resultErr = <-errorChan
		s.log.Error("[Album] service error: " + resultErr.Error())
		return models.Album{}, errors.New("[Album] service error: " + resultErr.Error())
	}

	return models.Album{
		ID:          preliminaryAlbum.ID,
		Name:        preliminaryAlbum.Name,
		Cover:       preliminaryAlbum.Cover,
		ReleaseDate: preliminaryAlbum.ReleaseDate,
		Artists:     <-artistsChan,
		Tracks:      <-tracksChan,
	}, nil
}
