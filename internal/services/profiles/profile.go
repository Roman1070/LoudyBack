package profiles

import (
	"context"
	"errors"
	albumsModels "loudy-back/internal/domain/models/albums"
	artistsModels "loudy-back/internal/domain/models/artists"
	models "loudy-back/internal/domain/models/profiles"
	tracksModels "loudy-back/internal/domain/models/tracks"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ProfilesService) Profile(ctx context.Context, id primitive.ObjectID) (models.Profile, error) {
	s.log.Info("[Profile] server started")

	profile, err := s.profiles.Profile(ctx, id)
	if err != nil {
		s.log.Error("[Profile] service error: " + err.Error())
		return models.Profile{}, errors.New("[Profile] service error: " + err.Error())
	}

	wg := sync.WaitGroup{}
	artistsChan := make(chan []artistsModels.ArtistLight, 1)
	albumsChan := make(chan []albumsModels.AlbumLight, 1)
	tracksChan := make(chan []tracksModels.Track, 1)
	errorChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		artistsLight, err := s.artistsProvider.ArtistsLight(ctx, profile.SavedArtistsIds)
		if err != nil {
			errorChan <- errors.New("[Profile] service error: " + err.Error())
			return
		}
		artistsChan <- artistsLight
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		albumsLight, err := s.albumsProvider.AlbumsLight(ctx, profile.SavedAlbumsIds)
		if err != nil {
			errorChan <- errors.New("[Profile] service error: " + err.Error())
			return
		}
		albumsChan <- albumsLight
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		tracks, err := s.tracksProvider.Tracks(ctx, profile.SavedTracksIds)
		if err != nil {
			errorChan <- errors.New("[Profile] service error: " + err.Error())
			return
		}
		tracksChan <- tracks
		wg.Done()
	}()

	wg.Wait()
	err = <-errorChan
	if err != nil {
		s.log.Error("[Profile] service error: " + err.Error())
		return models.Profile{}, errors.New("[Profile] service error: " + err.Error())
	}
	return models.Profile{
		ID:           profile.ID,
		Name:         profile.Name,
		Avatar:       profile.Avatar,
		Bio:          profile.Bio,
		SavedTracks:  <-tracksChan,
		SavedAlbums:  <-albumsChan,
		SavedArtists: <-artistsChan,
	}, nil
}
