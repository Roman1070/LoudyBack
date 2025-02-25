package tracks

import (
	"context"
	"errors"
	models "loudy-back/internal/domain/models/tracks"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *TracksService) Track(ctx context.Context, id primitive.ObjectID) (models.Track, error) {
	s.log.Info("[Track] service started")

	trackPreliminary, err := s.tracks.Track(ctx, id)
	if err != nil {
		s.log.Error("[Track] service error: " + err.Error())
		return models.Track{}, errors.New("[Track] service error: " + err.Error())
	}

	artists, err := s.artistsProvider.ArtistsLight(ctx, trackPreliminary.ArtistsIds)
	if err != nil {
		s.log.Error("[Track] service error: " + err.Error())
		return models.Track{}, errors.New("[Track] service error: " + err.Error())
	}

	return models.Track{
		ID:       trackPreliminary.ID,
		Name:     trackPreliminary.Name,
		Cover:    trackPreliminary.Cover,
		Duration: trackPreliminary.Duration,
		AlbumID:  trackPreliminary.AlbumID,
		Artists:  artists,
	}, nil
}
