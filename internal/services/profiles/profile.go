package profiles

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	albumsModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	models "loudy-back/internal/domain/models/profiles"
	tracksModels "loudy-back/internal/domain/models/tracks"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) Profile(ctx context.Context, id primitive.ObjectID) (models.Profile, error) {
	s.log.Info("[Profile] service started")

	profile, resultErr := s.profiles.Profile(ctx, id)
	if resultErr != nil {
		s.log.Error("[Profile] service error: " + resultErr.Error())
		return models.Profile{}, errors.New("[Profile] service error: " + resultErr.Error())
	}

	wg := sync.WaitGroup{}
	artistsChan := make(chan []artistsModels.ArtistLight, 1)
	albumsChan := make(chan []albumsModels.AlbumLight, 1)
	//tracksChan := make(chan []tracksModels.Track, 1)
	errorChan := make(chan error, 2)

	wg.Add(1)
	go func(ctx context.Context, artistsProvider ArtistsProvider, ids []primitive.ObjectID, log *slog.Logger) {
		log.Info("[Profile] go 2 started, provider = " + fmt.Sprint(artistsProvider) + " ids = " + fmt.Sprint(ids))
		artistsLight, err := artistsProvider.ArtistsLight(ctx, ids)
		if err != nil {
			errorChan <- errors.New("[Profile] service error: " + err.Error())
			return
		}
		s.log.Info("[Profile] go 1 pre-done")

		artistsChan <- artistsLight
		s.log.Info("[Profile] go 1 done")
		wg.Done()

	}(ctx, s.artistsProvider, profile.SavedArtistsIds, s.log)

	wg.Add(1)
	go func(ctx context.Context, albumsProvider AlbumsProvider, ids []primitive.ObjectID, log *slog.Logger) {
		log.Info("[Profile] go 2 started, provider = " + fmt.Sprint(albumsProvider) + " ids = " + fmt.Sprint(ids))
		albumsLight, err := albumsProvider.AlbumsLight(ctx, ids)
		if err != nil {
			errorChan <- errors.New("[Profile] service error: " + err.Error())
			return
		}
		log.Info("[Profile] go 2 pre-done")

		albumsChan <- albumsLight
		log.Info("[Profile] go 2 done")
		wg.Done()

	}(ctx, s.albumsProvider, profile.SavedAlbumsIds, s.log)
	/*
		wg.Add(1)
		go func(tracksProvider TracksProvider, ctx context.Context, ids []primitive.ObjectID, log *slog.Logger) {
			log.Info("[Profile] go 3 started, provider = " + fmt.Sprint(tracksProvider))
			tracks, err := tracksProvider.Tracks(ctx, ids)
			if err != nil {
				errorChan <- errors.New("[Profile] service error: " + err.Error())
				return
			}

			log.Info("[Profile] go 3 pre-done")
			tracksChan <- tracks

			log.Info("[Profile] go 3 done")
			wg.Done()
		}(s.tracksProvider, ctx, profile.SavedTracksIds, s.log)*/

	wg.Wait()
	s.log.Info("[Profile] Wait group finished waiting")

	errorsCount := len(errorChan)
	for i := 0; i < errorsCount; i++ {
		resultErr = <-errorChan
		s.log.Error("[Profile] service error: " + resultErr.Error())
		return models.Profile{}, errors.New("[Profile] service error: " + resultErr.Error())
	}

	//resultTracks := <-tracksChan
	resultAlbums := <-albumsChan
	resultArtists := <-artistsChan

	return models.Profile{
		ID:           profile.ID,
		Name:         profile.Name,
		Avatar:       profile.Avatar,
		Bio:          profile.Bio,
		SavedTracks:  []tracksModels.Track{},
		SavedAlbums:  resultAlbums,
		SavedArtists: resultArtists,
	}, nil
}
